//加好友申请
function ToFriendUserById(id) {
    $.ajax({
        type: "POST",
        url: httpUrl + '/friend/crete',
        data: {user_id: id},
        dataType: 'json',
        success: function (data) {
            if (data == 2) {
                swal("已经是好友关系请勿发送", "", "error", {
                    button: "继续"
                })
            } else {
                swal("已发送好友请求", "", "success", {
                    button: "继续"
                })
            }
        },
    });
}

//切换聊天对象的时候更新当前将要聊天对象的user_Id
function change_user(obj, un = 0) {
    $(".chat").children('.index-chat').remove()
    $("#bottom-show").show()
    $("#scroll").show()
    //点开表示读消息 对应未读数减去
    sum = $("#un-message").text()
    if (sum <= 0) {
        $("#un-message").html(0)
    } else {
        current_count = sum - un
        $("#un-message").html(current_count)
    }

    currentId = $("#current_id").val();
    var from_id = $(obj).parents("li").find("#from_id").val();
    var to_id = $(obj).parents("li").find("#to_id").val();
    var to_name = $(obj).parents("li").find("#to_name").val();
    //未读消息清空
    $("#unread_count" + to_id).text("未读消息:0")
    $("#current_name").text(to_name)
    sendId = currentId == from_id ? to_id : from_id;

    //查看消息记录
    $.ajax({
        type: "POST",
        url: httpUrl + '/chat/chatAll',//被请求的API接口地址
        data: {to_id: sendId},
        dataType: 'json',
        success: function (data) {
            $(".chat-child").remove()
            if (data == null) {
                return
            }

            for (var i = 0; i < data.length; i++) {

                type = data[i].FromId == currentId ? 1 : 2
                content = replace_em(data[i].Content);
                if (data[i].ContentType == 2) {
                    content = '<img src="'+resourceUrl + data[i].Content + '" width="100px"  height="100px" />'
                }
                createChatHtmlNode(content, data[i].SendTime, type)
            }
            msgScrollTOp()
        },
    });
}

//获取好友列表
function getFriends() {
    $.ajax({
        type: "GET",
        url: httpUrl + '/friends',
        dataType: 'json',
        success: function (data) {
            $("#friends-list").children().remove()
            if (data == null) {
                return
            }
            for (i = 0; i < data.length; i++) {
                var friend_html = '<li>\n' +
                    '                                        <input type="hidden" id="from_id" name="from_id" value="' + data[i].UserId + '"><br>\n' +
                    '                                        <input type="hidden" id="to_id" name="to_id" value="' + data[i].FriendId + '"><br>\n' +
                    '                                        <input type="hidden" id="to_name" name="to_name" value="' + data[i].FriendName + '"><br>\n' +
                    '                                        <a href="#chat1" onclick="change_user(this,' + data[i].UnreadMessage + ')">\n' +
                    '                                            <div class="status online"><img\n' +
                    '                                                        src="/static/dist/img/avatars/avatar-female-3.jpg" alt="avatar"><i\n' +
                    '                                                        data-eva="radio-button-on"></i></div>\n' +
                    '                                            <div class="content">\n' +
                    '                                                <h5>' + data[i].FriendName + '</h5>\n' +
                    '                                                <span>中国</span>\n' +
                    '                                            </div>\n' +
                    '                                            <div class="icon"><i data-eva="person"></i></div>\n' +
                    '                                        </a>\n' +
                    '                                    </li>'

                $("#friends-list").append(friend_html)
            }
        },
    });
}


//获取好友申请
function GetApplications() {
    $.ajax({
        type: "GET",
        url: httpUrl + '/user/apps',//被请求的API接口地址
        dataType: 'json',
        success: function (data) {
            $("#app-counts").children(".sum_notes").html('')
            $("#system_notices").children().remove()
            if (data == null) {
                $("#app-counts").children(".sum_notes").append(0)
                return
            }

            $("#app-counts").children(".sum_notes").append(data.length)
            for (i = 0; i < data.length; i++) {
                addHandleFriendHtml(data[i].FromId, data[i].Remark)
            }
        },
    });
}
