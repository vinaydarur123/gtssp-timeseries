package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/your-org/gtssp/internal/util"
)

func StartOutputServer() {

	// Clear traces
	http.HandleFunc("/clear", func(w http.ResponseWriter, r *http.Request) {
		util.ClearTraces()
		http.Redirect(w, r, "/output", http.StatusSeeOther)
	})

	// Export traces as JSON
	http.HandleFunc("/export", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(util.GetTraces())
	})

	// Output UI
	http.HandleFunc("/output", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		mode := r.URL.Query().Get("mode")
		isResults := mode == "results"

		traces := util.GetTraces()

		limit := len(traces)
		if v := r.URL.Query().Get("limit"); v != "" {
			if n, err := strconv.Atoi(v); err == nil && n < limit {
				limit = n
			}
		}

		fmt.Fprintln(w, `
<!DOCTYPE html>
<html>
<head>
<title>GTSPP Output</title>
<style>
body { font-family: Arial; background:#f4f6f8; padding:20px }
.stage { background:#fff; padding:15px; margin-bottom:20px; border-radius:8px }
.noop { color:#888 }
.changed { color:#1b7f3a; font-weight:bold }
pre { background:#eee; padding:10px }
.controls { margin-bottom:20px }
.results body { background:white }
</style>

<script>
let refreshInterval=null;

function toggleRefresh(){
	const on=document.getElementById("autorefresh").checked;
	localStorage.setItem("autorefresh",on);
	if(on) refreshInterval=setInterval(()=>location.reload(),5000);
	else clearInterval(refreshInterval);
}

window.onload=function(){
	const saved=localStorage.getItem("autorefresh")==="true";
	const cb=document.getElementById("autorefresh");
	if(cb){cb.checked=saved; if(saved) toggleRefresh();}
}
</script>
</head>
<body class="`)

		if isResults {
			fmt.Fprint(w, "results")
		}

		fmt.Fprintln(w, `">

<h1>GTSPP Metric Output</h1>
<p>Live visualization of metric validation, transformation, and storage</p>`)

		if !isResults {
			fmt.Fprintln(w, `
<div class="controls">
<form method="POST" action="/clear" style="display:inline;">
<button>🧹 Clear Traces</button>
</form>

<label style="margin-left:20px;">
<input type="checkbox" id="autorefresh" onchange="toggleRefresh()"> 🔄 Auto Refresh
</label>

<a href="/export" target="_blank" style="margin-left:20px;">⬇ Export JSON</a>
</div>

<label>Timeline:
<input type="range" min="1" max="`+strconv.Itoa(len(traces))+`" value="`+strconv.Itoa(limit)+`"
onchange="location='?limit='+this.value">
</label>
<hr>
`)
		}

		start := 0
		if len(traces) > limit {
			start = len(traces) - limit
		}
		for _, t := range traces[start:] {
			fmt.Fprintf(w, `<div class="stage"><h3>%s</h3>`, t.Stage)

			if t.NoOp {
				fmt.Fprintf(w, `<p class="noop">🟡 %s</p>`, t.Summary)
			} else {
				fmt.Fprintf(w, `<p class="changed">🟢 %s</p>`, t.Summary)
			}

			fmt.Fprintf(w,
				`<pre>Name: %s
Value: %v
Timestamp: %v
Labels: %+v</pre>`,
				t.Metric.Name,
				t.Metric.Value,
				t.Metric.Timestamp,
				t.Metric.Labels)

			fmt.Fprintln(w, `</div>`)
		}

		fmt.Fprintln(w, `
</body>
</html>`)
	})

	go func() {
		fmt.Println("🔗 Output available at: http://localhost:8082/output")
		http.ListenAndServe(":8082", nil)
	}()
}
