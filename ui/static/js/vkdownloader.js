function openMedia(id) {
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/open-media/" + id, true);
    console.log("open-media " , id);
    xhr.send();
}

function onCheckMedia(event, id) {
	console.log("onCheckMedia " , id);
	let checkbox = document.getElementById(id);
    let parentDiv = checkbox.parentNode.parentNode;
    if(checkbox.checked) {
        if (event.shiftKey)
            parentDiv.style.backgroundColor = '#5555FF';
        else
            parentDiv.style.backgroundColor = '#FFA500';
    } else {
       parentDiv.style.backgroundColor = 'white';
    }
}

function openURL(event) {
    event.preventDefault();
    let appId = document.getElementById("app_id").value;
    console.log("openURL: " + appId);
    let appUrl = "https://oauth.vk.com/authorize?client_id="
        + appId + "&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=notify,friends,photos,audio,video,docs,notes,pages,status,wall,groups,notifications&response_type=token&v=5.131"
    // open in new page
    window.open(appUrl, "_blank");
    // prepare AJAX
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/set-app-id", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    // handler
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            //alert(xhr.responseText);
            location.reload()
        }
    };

    xhr.send("app_id=" + appId);
}

function postURL(event) {
    event.preventDefault();
    let appUrl = document.getElementById("app_url").value;
    console.log("postURL: " + appUrl);

    // window.open("https://google.com", "_blank");
    // prepare AJAX
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/set-app-url", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    // handler
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            //alert(xhr.responseText);
            location.reload()
        }
    };
    // send
    let prm = "app_url=" + appUrl;
    console.log("send to: " + prm);
    xhr.send(prm);
}

function setProxy(event) {
    event.preventDefault();
    let proxyUrl = document.getElementById("proxy_addr").value;
    let proxyUse = document.getElementById("proxy_use").checked;
    console.log("proxyURL: " + proxyUrl + " , proxyUse: " + proxyUse);
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/set-proxy", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    // handler
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            //alert(xhr.responseText);
            location.reload()
        }
    };
    // send
    let prm = "proxy_url=" + proxyUrl + "&proxy_use=" + proxyUse;
    console.log("proxy prm: " + prm);
    xhr.send(prm);
}

function getIP(event) {
	event.preventDefault();
	console.log("GetIP");
	let xhr = new XMLHttpRequest();
    xhr.open("POST", "/get-ip", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
			console.log(xhr.responseText)
            location.reload()
        }
    };
    xhr.send("");
}

function openTestURL(ibase) {
    let base = document.getElementById(ibase).textContent;
    console.log(base)
    let token = document.getElementById('token').textContent;
    console.log(token)
    let url = base + token;
    window.open(url, "_blank");
}

function setSearch(extraArgs) {
	console.log("setSearch ", extraArgs);
   // Get the name from the form
   let text = document.getElementById('search').value;
   let currentUrl = window.location.href.split('?')[0];
   window.location.href = currentUrl + '?' +pkArgs('search', encodeURIComponent(text), extraArgs);
}

function clearSearch(extraArgs) {
    let currentUrl = window.location.href.split('?')[0];
	window.location.href = currentUrl + '?' + extraArgs;
}

function applyFilters(extraArgs) {
    let currentUrl = window.location.href.split('?')[0];
    let code = getChk('f_my') +  getChk('f_bm') + getChk('f_fr') + getChk('f_gr') + getChk('f_lk') + getChk('f_cm')
    + getChk('f_de') + getChk('f_ba') + getChk('f_gb');
    if(filtersIsEmpty(code))
		window.location.href = currentUrl + '?' +extraArgs;
	else
		window.location.href = currentUrl + '?' +pkArgs('filters', code, extraArgs);
}

function setTags(event, mode) {
    event.preventDefault();
    console.log("setTags: mode=", mode);
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/set-tags", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            console.log(xhr.responseText)
            location.reload()
        }
    };
    const tag = document.getElementById('newtag').value;
    if(!isValidTag(tag)) {
        alert("incorrect new tag");
        return;
    }
    document.getElementById('newtag').value = ""
    const checks  = getCheckboxes(".chk_id:checked", "~");
    const settags = getCheckboxes(".chk_tag:checked", "~");
    const clrtags = getCheckboxes(".chk_tag:indeterminate", "~");
    const prm = pkArgs("newtag", tag, "settags", settags, "clrtags", clrtags, "checkedids", checks, "mode", mode);
    console.log(prm)
    xhr.send(prm);
}

function applyTags(extraArgs) {
    let currentUrl = window.location.href.split('?')[0];
    let settags = getCheckboxes(".chk_tag:checked", "~");
    let clrtags = getCheckboxes(".chk_tag:indeterminate", "-");
    let tags = "";
    if(settags != "") {
		tags += "~";
		tags += settags;
	}
	if(clrtags != "") {
		tags += "-";
		tags += clrtags;
	}
    console.log("TAGS: ", tags);
    if(tags == "")
        window.location.href = currentUrl + '?' + extraArgs;
    else
        window.location.href = currentUrl + '?' + pkArgs('tags', tags, extraArgs);
}

function loadTags(str) {
    console.log("loadTags:", str)
    let arr = str.split(/([-~])/);
    let len = arr.length;
    console.log(arr, len)
    for(let i=1; i<len; i+=2) {
        console.log("loadTags word: ", arr[i], arr[i+1])
        if (arr[i]=='~') {
            setChk(arr[i+1], '1')
        } else if (arr[i]=='-') {
            setChk(arr[i+1], '2');
        }
    }
}

function isValidTag(str) {
    return /^[0-9A-Z_a-z]*$/.test(str);
}

//
function exploreActor(vkurl) {
    console.log("explore: ", vkurl)
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/explore-actor", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onload = function() {
        if (this.status >= 200 && this.status < 400) {
            let newUrl = this.responseText;
            console.log("OnLoad:", newUrl)
            window.location.href = newUrl;
        }
    };

    const data = "id=" + vkurl;
    xhr.send(data);
}

function addActor(vkid) {
    console.log("addActor: ", vkid)
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/add-actor", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onload = function() {
        if (this.status >= 200 && this.status < 400) {
            let newUrl = this.responseText;
            console.log("OnLoad:", newUrl)
            window.location.href = newUrl;
        }
    };

    const data = "id=" + vkid;
    xhr.send(data);
}

// HELPERS FOR UPDATE DB QUERIES

function sendUpdateQuery(url) {
    // prepare AJAX
    let xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            location.reload()
        }
    };

    console.log(url);
    xhr.send();
}

function sendUpdateQueryARG(url, argvalue) {
    console.log("sendUpdateQueryARG: ", argvalue)
    let xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            location.reload()
        }
    };

    const data = "id=" + argvalue;
    console.log(url);
    xhr.send(data);
}

function getCheckboxes(selector, delimiter) {
    const checkboxes = document.querySelectorAll(selector);
    const checkboxValues = [];
    for (let i = 0; i < checkboxes.length; i++) {
        checkboxValues.push(checkboxes[i].id);
    }
    console.log(checkboxValues)
    const data = encodeURIComponent(checkboxValues.join(delimiter));
    return data;
}

function sendUpdateQueryCB(url) {
    // prepare AJAX
    let xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            location.reload()
        }
    };

    const data = getCheckboxes(".chk_id:checked", "~");
    const prm = pkArgs("checkedids", data);
    xhr.send(prm);
    console.log(url);
 }

function checkAll()
{
    const mainCheckbox = document.getElementById('all');
    const checkboxes = document.querySelectorAll('.chk_id');
    checkboxes.forEach(function(checkbox) {
       checkbox.checked = mainCheckbox.checked;
    });
}

function ts(cb) {
    if (cb.readOnly) cb.checked=cb.readOnly=false;
    else if (!cb.checked) cb.readOnly=cb.indeterminate=true;
}

function getChk(id) {
    let cb = document.getElementById(id);
    if(cb.indeterminate) return '2';
    if(cb.checked) return '1';
    return '0';
}

function setChk(id, st) {
    let cb = document.getElementById(id);
    if(st=='2') {
        cb.readOnly=cb.indeterminate=true;
    } else if(st=='1') {
        cb.readOnly=false;
        cb.checked=true;
    } else if(st=='0') {
        cb.checked=cb.readOnly=false;
    }
}

function filtersIsEmpty(str) {
    for (let i = 0; i<str.length; i++)
        if(str[i]!='0')
            return false;
    return true;
}

function pkArgs() {
    // variadic function
 //   console.log("pkArgs: ", arguments.length);
    // converts args to string like "arg1=arg2&arg3=arg4&arg5=arg6&arg7"
    let i = 0;
    let res = '';
    // pairs
    for (i = 0; i < arguments.length-1; i+=2) {
//		console.log("pair: ", arguments[i], " == ", arguments[i + 1]);
        if(arguments[i+1] != '') {			
            if(i>0)
                res += '&';
            res += arguments[i];
            res += '=';
            res += arguments[i + 1];
        }
    }
    // trailing argument
    if(i < arguments.length && arguments[i]!="") {
//		console.log("Trailing arg: ", arguments[i]);
        if(i>0 && arguments[i][0]!="&")
            res += '&';
        res += arguments[i];
    }
    return res;
}


// GLOBAL AREA

document.addEventListener("DOMContentLoaded", function(event) {
    console.log("init page")

    let statusbar = document.getElementById('statusbar');

    // Update the status in the HTML
    let source = new EventSource("http://localhost:8080/get-server-status");
    source.onmessage = function(event) {
        statusbar.textContent = event.data;
    };

    //
    let started = false;
    let btn = document.getElementById("controlbtn");

    // Get the initial state of the goroutine
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/get-worker-status");
    xhr.onload = function() {
        if (xhr.status === 200) {
            started = (xhr.responseText === "true");
            btn.textContent = started ? "STOP" : "START";
        }
    };
    xhr.send();

    btn.addEventListener("click", function(event) {
        event.preventDefault();
        if (started) {
            console.log("stop...")
            btn.textContent = "STOPPING...";
            // Stop the goroutine
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "http://127.0.0.1:8080/stop-worker");
            xhr.onload = function() {
                console.log("Worker stopped.");
                started = false;
                btn.textContent = "START";
            };
            xhr.onerror = function() {
                console.log("Error stopping worker.");
            };
            xhr.send();
        } else {
            console.log("start...")
            btn.textContent = "STARTING...";
            // Start the goroutine
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "http://127.0.0.1:8080/start-worker");
            xhr.onload = function() {
                console.log("Worker started.");
                started = true;
                btn.textContent = "STOP";
            };
            xhr.onerror = function() {
                console.log("Error starting worker.");
            };
            xhr.send();
        }
    });
});
