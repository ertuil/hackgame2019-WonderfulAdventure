
function addMsg(from, content) {
    if ( typeof(content) == "undefined") {
        return
    }
    let alist = $("#alist");
    alist.append("<h4>"+from+":</h4>");
    alist.append("<p>"+content+"</p>");
}

function updateState(state) {
    $("#info_name").text(state["Name"]);
    $("#info_money").text(state["Money"]);
    $("#info_attack").text(state["Attack"]);
}

function loadOptions(options) {
    opts = options;
    let st = "";
    for(let i = 0, len = options.length; i < len; i++){
        st += "<option value =\""+ i +"\">"+options[i]+"</option>"
    }
    $("#input").html(st);
}