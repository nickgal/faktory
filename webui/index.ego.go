// Generated by ego.
// DO NOT EDIT

//line index.ego:1

package webui

import "fmt"
import "html"
import "io"
import "context"

import (
	"github.com/contribsys/faktory/client"
	"github.com/contribsys/faktory/util"
	"net/http"
)

func ego_index(w io.Writer, req *http.Request) {
	ego_layout(w, req, func() {
//line index.ego:12
		_, _ = io.WriteString(w, "\n<script type=\"text/javascript\" src=\"")
//line index.ego:12
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(relative(req, "/static/dashboard.js"))))
//line index.ego:12
		_, _ = io.WriteString(w, "\"></script>\n<div class= \"dashboard clearfix\">\n  <h3 >\n    ")
//line index.ego:15
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Dashboard"))))
//line index.ego:16
		_, _ = io.WriteString(w, "\n    <span class=\"beacon\">\n      <span class=\"ring\"></span>\n      <span class=\"dot\"></span>\n    </span>\n  </h3>\n  <div class=\"interval-slider ltr\">\n    <span class=\"interval-slider-label\">")
//line index.ego:22
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "PollingInterval"))))
//line index.ego:22
		_, _ = io.WriteString(w, ":</span>\n    <span class=\"current-interval\">5 sec</span>\n    <br/>\n    <input type=\"range\" min=\"2000\" max=\"20000\" step=\"1000\" value=\"5000\"/>\n  </div>\n</div>\n\n<div class=\"row chart\">\n  <div id=\"realtime\" data-processed-label=\"")
//line index.ego:30
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Processed"))))
//line index.ego:30
		_, _ = io.WriteString(w, "\" data-failed-label=\"")
//line index.ego:30
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Failed"))))
//line index.ego:30
		_, _ = io.WriteString(w, "\"></div>\n  <div id=\"realtime-legend\"></div>\n</div>\n\n<div class=\"row chart\">\n  <h5>\n    <span class=\"history-heading\">")
//line index.ego:36
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "History"))))
//line index.ego:36
		_, _ = io.WriteString(w, "</span>\n    <a href=\"/?days=7\" class=\"history-graph ")
//line index.ego:37
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(daysMatches(req, "7", false))))
//line index.ego:37
		_, _ = io.WriteString(w, "\">")
//line index.ego:37
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "OneWeek"))))
//line index.ego:37
		_, _ = io.WriteString(w, "</a>\n    <a href=\"/\" class=\"history-graph ")
//line index.ego:38
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(daysMatches(req, "30", true))))
//line index.ego:38
		_, _ = io.WriteString(w, "\">")
//line index.ego:38
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "OneMonth"))))
//line index.ego:38
		_, _ = io.WriteString(w, "</a>\n    <a href=\"/?days=90\" class=\"history-graph ")
//line index.ego:39
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(daysMatches(req, "90", false))))
//line index.ego:39
		_, _ = io.WriteString(w, "\">")
//line index.ego:39
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "ThreeMonths"))))
//line index.ego:39
		_, _ = io.WriteString(w, "</a>\n    <a href=\"/?days=180\" class=\"history-graph ")
//line index.ego:40
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(daysMatches(req, "180", false))))
//line index.ego:40
		_, _ = io.WriteString(w, "\">")
//line index.ego:40
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "SixMonths"))))
//line index.ego:40
		_, _ = io.WriteString(w, "</a>\n  </h5>\n\n  <div id=\"history\" data-processed-label=\"")
//line index.ego:43
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Processed"))))
//line index.ego:43
		_, _ = io.WriteString(w, "\" data-failed-label=\"")
//line index.ego:43
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Failed"))))
//line index.ego:43
		_, _ = io.WriteString(w, "\" data-processed=\"")
//line index.ego:43
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(processedHistory(req))))
//line index.ego:43
		_, _ = io.WriteString(w, "\" data-failed=\"")
//line index.ego:43
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(failedHistory(req))))
//line index.ego:43
		_, _ = io.WriteString(w, "\" data-update-url=\"/stats\"></div>\n  <div id=\"history-legend\"></div>\n</div>\n\n<br/>\n<h5>Faktory</h5>\n<div class=\"faktory-wrapper\">\n  <div class=\"stats-container\">\n    <div class=\"stat\">\n      <h3 class=\"faktory_version\">")
//line index.ego:52
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(client.Version)))
//line index.ego:52
		_, _ = io.WriteString(w, "</h3>\n      <p>")
//line index.ego:53
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Version"))))
//line index.ego:53
		_, _ = io.WriteString(w, "</p>\n    </div>\n\n    <div class=\"stat\">\n      <h3 class=\"uptime_in_days\">")
//line index.ego:57
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(uptimeInDays(req))))
//line index.ego:57
		_, _ = io.WriteString(w, "</h3>\n      <p>")
//line index.ego:58
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Uptime"))))
//line index.ego:58
		_, _ = io.WriteString(w, "</p>\n    </div>\n\n    <div class=\"stat\">\n      <h3 class=\"connections\">")
//line index.ego:62
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(ctx(req).Server().Stats.Connections)))
//line index.ego:62
		_, _ = io.WriteString(w, "</h3>\n      <p>")
//line index.ego:63
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Connections"))))
//line index.ego:63
		_, _ = io.WriteString(w, "</p>\n    </div>\n\n    <div class=\"stat\">\n      <h3 class=\"used_memory_human\">")
//line index.ego:67
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(util.MemoryUsage())))
//line index.ego:67
		_, _ = io.WriteString(w, "</h3>\n      <p>")
//line index.ego:68
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "MemoryUsage"))))
//line index.ego:68
		_, _ = io.WriteString(w, "</p>\n    </div>\n\n    <div class=\"stat\">\n      <h3 class=\"commands\">")
//line index.ego:72
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(ctx(req).Server().Stats.Commands)))
//line index.ego:72
		_, _ = io.WriteString(w, "</h3>\n      <p>")
//line index.ego:73
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "CommandsExecuted"))))
//line index.ego:73
		_, _ = io.WriteString(w, "</p>\n    </div>\n  </div>\n</div>\n  ")
//line index.ego:77
	})
//line index.ego:78
	_, _ = io.WriteString(w, "\n")
//line index.ego:78
	// vim: set ft=html
//line index.ego:79
	_, _ = io.WriteString(w, "\n")
//line index.ego:79
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString
