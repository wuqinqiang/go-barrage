//客户端也处理一些xss
function stripscript(s) {
    var pattern = new RegExp("[%--`~!@#$^&*()=|{}':;',\\.<>/?~！@#￥……&*（）——|{}【】‘；：”“'。，、？]")        //格式 RegExp("[在中间定义特殊过滤字符]")
    var rs = "";
    for (var i = 0; i < s.length; i++) {
        rs = rs + s.substr(i, 1).replace(pattern, '');
    }
    return rs;
}

//替换成表情展示信息
function replace_em(str) {

    str = str.replace(/\</g, '&lt;');

    str = str.replace(/\>/g, '&gt;');

    str = str.replace(/\n/g, '<br/>');

    str = str.replace(/\[em_([0-9]*)\]/g, '<img src="http://localhost:8080/static/emm/arclist/$1.gif" border="0" />');
    return str;

}

function addEmmHtml(index,content) {
    content=replace_em(content)
    $("last-message"+index).html(content)
}


//消息聊天到底最底部
function msgScrollTOp() {
    var div = document.getElementById('scroll');
    div.scrollTop = div.scrollHeight
}
