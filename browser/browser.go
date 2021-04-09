package browser

import (
	"html/template"
	"net/http"
)

func LogSocketViewHandler(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+r.URL.Path+"/ws")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
</head>
<body>
<center>
<input type="text" id="search" onkeyup="filterTable()" placeholder="Filter...">
<table id="logHeaders" style="text-align:left; width:80%;" >
<tbody>
<tr class="header">
<th style="width:20%;">TimeStamp</th>
<th style="width:5%;">Level</th>
<th style="width:65%;">Output</th>
<th style="width:10%;">Source</th>
</tr>
</tbody>
</table>

<div id="tableWrapper">
<table id="logs" style="text-align:left; width:100%;" >
<tbody id="tbodylogs">
<tr class="header">
<th style="width:20%;"></th>
<th style="width:5%;"></th>
<th style="width:65%;"></th>
<th style="width:10%;"></th>
</tr>

</tbody>
</table>
</div>
<br>
<input class="button" type="button" id="download" value="Download Logs" style="background-color:#3f51b5;"/>
<input class="button" type="button" id="delete" value="Delete Logs" style="background-color:#f44336"/>
</center>
</body>
<footer>
<style>
#tableWrapper{
	overflow-y: scroll;
	display: flow-root;
	width: 80%;
	height: 80vh;
}
td,tr{
	height: min-content;
}
.button{
	display: inline-block;
	width: 5vw;
	height: 5vh;
}
</style>
<script>  
var logTable = document.getElementById("logs");
var logTableB = document.getElementById("tbodylogs");
var ws = null;
var application = "demo-commit"
var logs = [];
function sleep(ms) {
	return new Promise(resolve => setTimeout(resolve, ms));
}
function download(filename, text) {
	var element = document.createElement('a');
	element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
	element.setAttribute('download', filename);
	element.style.display = 'none';
	document.body.appendChild(element);
	element.click();
	document.body.removeChild(element);
}
function openSocket() {
	if (ws) {
		return false;
	}
	ws = new WebSocket("{{.}}");
	ws.onclose = async function(evt) {
		ws = null;
		while(ws == null){
			openSocket()
			await sleep(5000);
		}
	}
	ws.onmessage = function(evt) {
		var entry = JSON.parse(evt.data)
		logs.push(entry)
		var row = document.createElement('tr');
		var ts = document.createElement('td');
		var tst = document.createTextNode(entry.timestamp);
		ts.appendChild(tst);
		row.appendChild(ts);
		var ts = document.createElement('td');
		var tst = document.createTextNode(entry.level);
		ts.appendChild(tst);
		row.appendChild(ts);
		var ts = document.createElement('td');
		var tst = document.createTextNode(entry.output);
		ts.appendChild(tst);
		row.appendChild(ts);
		var ts = document.createElement('td');
		var tst = document.createTextNode(entry.file);
		ts.appendChild(tst);
		row.appendChild(ts);
		var bg="";
		switch(entry.level){
		case "INFO":
			bg="white";
			break;
		case "ERROR":
			bg="#f44336";
			break;
		case "WARN":
			bg="#fb8c00"
			break;
		case "TRACE":
			bg="#E1F5FE"
			break;
		case "DEBUG":
			bg="#B3E5FC"
			break;
		default:
			bg="white"
			break;
		}
		row.style.backgroundColor=bg
		logTableB.append(row)
		filterTable()
	}
	ws.onerror = function(evt) {
		if (evt != null && evt.data != null){
			// handle error here
		}
	}
}
function clearTable(){
	if(!window.confirm("Are you sure you want to delete all logs?")){
		return
	}
	logs = []

	while (logTableB.childNodes.length > 1) {
		logTableB.removeChild(logTableB.childNodes[1]);
	}
}
function filterTable() {
	var cols, input, filter, table, tr, td, i, txtValue, w;
	input = document.getElementById("search");
	filter = input.value;
	table = logTableB;
	tr = table.getElementsByTagName("tr");
	for (i = 1; i < tr.length; i++) {
		cols = tr[i].getElementsByTagName("td");
		var visible = false;
		for (w = 0; w < cols.length; w++){
			if (!visible && cols[w]) {
				td = cols[w]
				txtValue = td.textContent || td.innerText;
				if (txtValue.indexOf(filter) > -1) {
					visible = true
				} 
			}
		}
		if(visible){

			tr[i].style.display = "";
		} else {

			tr[i].style.display = "none";
		}

	}
}
document.getElementById("delete").addEventListener("click", function(){

	clearTable()
}, false);
document.getElementById("download").addEventListener("click", function(){

	download(application+'.json',JSON.stringify(logs));
}, false);
openSocket();
</script>

</footer>
</html>
`))
