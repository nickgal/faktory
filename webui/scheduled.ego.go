// Generated by ego.
// DO NOT EDIT

//line scheduled.ego:1

package webui

import "fmt"
import "html"
import "io"
import "context"

import (
	"net/http"

	"github.com/contribsys/faktory/client"
	"github.com/contribsys/faktory/storage"
)

func ego_listScheduled(w io.Writer, req *http.Request, set storage.SortedSet, count, currentPage uint64) {
	totalSize := uint64(set.Size())

//line scheduled.ego:14
	_, _ = io.WriteString(w, "\n\n")
//line scheduled.ego:15
	ego_layout(w, req, func() {
//line scheduled.ego:16
		_, _ = io.WriteString(w, "\n\n<header class=\"row\">\n  <div class=\"col-sm-5\">\n    <h3>")
//line scheduled.ego:19
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "ScheduledJobs"))))
//line scheduled.ego:19
		_, _ = io.WriteString(w, "</h3>\n  </div>\n  ")
//line scheduled.ego:21
		if totalSize > 0 && totalSize > count {
//line scheduled.ego:22
			_, _ = io.WriteString(w, "\n    <div class=\"col-sm-4\">\n      ")
//line scheduled.ego:23
			ego_paging(w, req, "/scheduled", totalSize, count, currentPage)
//line scheduled.ego:24
			_, _ = io.WriteString(w, "\n    </div>\n  ")
//line scheduled.ego:25
		}
//line scheduled.ego:26
		_, _ = io.WriteString(w, "\n  ")
//line scheduled.ego:26
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(filtering("scheduled"))))
//line scheduled.ego:27
		_, _ = io.WriteString(w, "\n</header>\n\n")
//line scheduled.ego:29
		if totalSize > 0 {
//line scheduled.ego:30
			_, _ = io.WriteString(w, "\n\n  <form action=\"")
//line scheduled.ego:31
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(root(req))))
//line scheduled.ego:31
			_, _ = io.WriteString(w, "/scheduled\" method=\"post\">\n    ")
//line scheduled.ego:32
			_, _ = fmt.Fprint(w, csrfTag(req))
//line scheduled.ego:33
			_, _ = io.WriteString(w, "\n    <div class=\"table_container\">\n      <table class=\"table table-striped table-bordered table-white\">\n        <thead>\n          <tr>\n            <th class=\"checkbox-column\">\n              <input type=\"checkbox\" class=\"check_all\" />\n            </th>\n            <th>")
//line scheduled.ego:40
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "When"))))
//line scheduled.ego:40
			_, _ = io.WriteString(w, "</th>\n            <th>")
//line scheduled.ego:41
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Queue"))))
//line scheduled.ego:41
			_, _ = io.WriteString(w, "</th>\n            <th>")
//line scheduled.ego:42
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Job"))))
//line scheduled.ego:42
			_, _ = io.WriteString(w, "</th>\n            <th>")
//line scheduled.ego:43
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Arguments"))))
//line scheduled.ego:43
			_, _ = io.WriteString(w, "</th>\n          </tr>\n        </thead>\n        ")
//line scheduled.ego:46
			setJobs(set, count, currentPage, func(idx int, key []byte, job *client.Job) {
//line scheduled.ego:47
				_, _ = io.WriteString(w, "\n          <tr>\n            <td>\n              <input type=\"checkbox\" name=\"key\" value=\"")
//line scheduled.ego:49
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(string(key))))
//line scheduled.ego:49
				_, _ = io.WriteString(w, "\" />\n            </td>\n            <td>\n               <a href=\"")
//line scheduled.ego:52
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(root(req))))
//line scheduled.ego:52
				_, _ = io.WriteString(w, "/scheduled/")
//line scheduled.ego:52
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(string(key))))
//line scheduled.ego:52
				_, _ = io.WriteString(w, "\">")
//line scheduled.ego:52
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(relativeTime(job.At))))
//line scheduled.ego:52
				_, _ = io.WriteString(w, "</a>\n            </td>\n            <td>\n              <a href=\"")
//line scheduled.ego:55
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(root(req))))
//line scheduled.ego:55
				_, _ = io.WriteString(w, "/queues/")
//line scheduled.ego:55
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(job.Queue)))
//line scheduled.ego:55
				_, _ = io.WriteString(w, "\">")
//line scheduled.ego:55
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(job.Queue)))
//line scheduled.ego:55
				_, _ = io.WriteString(w, "</a>\n            </td>\n            <td><code>")
//line scheduled.ego:57
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(job.Type)))
//line scheduled.ego:57
				_, _ = io.WriteString(w, "</code></td>\n            <td>\n               <div class=\"args\">")
//line scheduled.ego:59
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(displayArgs(job.Args))))
//line scheduled.ego:59
				_, _ = io.WriteString(w, "</div>\n            </td>\n          </tr>\n        ")
//line scheduled.ego:62
			})
//line scheduled.ego:63
			_, _ = io.WriteString(w, "\n      </table>\n    </div>\n    <div class=\"pull-right flip\">\n      <button class=\"btn btn-primary\" type=\"submit\" name=\"action\" value=\"add_to_queue\">")
//line scheduled.ego:66
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "AddToQueue"))))
//line scheduled.ego:66
			_, _ = io.WriteString(w, "</button>\n      <button class=\"btn btn-danger\" type=\"submit\" name=\"action\" value=\"delete\">")
//line scheduled.ego:67
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Delete"))))
//line scheduled.ego:67
			_, _ = io.WriteString(w, "</button>\n    </div>\n  </form>\n")
//line scheduled.ego:70
		} else {
//line scheduled.ego:71
			_, _ = io.WriteString(w, "\n  <div class=\"alert alert-success\">")
//line scheduled.ego:71
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "NoScheduledFound"))))
//line scheduled.ego:71
			_, _ = io.WriteString(w, "</div>\n")
//line scheduled.ego:72
		}
//line scheduled.ego:73
		_, _ = io.WriteString(w, "\n")
//line scheduled.ego:73
	})
//line scheduled.ego:74
	_, _ = io.WriteString(w, "\n")
//line scheduled.ego:74
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString
