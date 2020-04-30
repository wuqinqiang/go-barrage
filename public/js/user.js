//获取用户列表
function getUsers() {
    $.ajax({
        type: "GET",
        url: httpUrl+'/users',
        dataType: 'json',
        success: function (data) {
            $("#userList").children().remove()
            for (var i = 0; i < data.length; i++) {
                userNode = '                                <li>\n' +
                    '                                    <div class="status online"><img src="/static/dist/img/avatars/avatar-male-6.jpg"\n' +
                    '                                                                    alt="avatar"><i data-eva="radio-button-on"></i>\n' +
                    '                                    </div>\n' +
                    '                                    <div class="content">\n' +
                    '                                        <h5>' + data[i].Name + '</h5>\n' +
                    '                                        <span>' + data[i].Email + '</span>\n' +
                    '                                    </div>\n' +
                    '                                    <div class="custom-control">\n' +
                    ' <button  onclick="ToFriendUserById(' + data[i].Id + ')" type="button" class="btn" data-toggle="modal"\n' +
                    '                                                data-target="#compose"><i\n' +
                    '                                                    data-eva="person-add" data-eva-animation="pulse">请求验证</i></button>'
                // '                                        <input type="button" onclick="ToFriendUserById(' + data[i].Id + ')" class="custom-control-input" id="user' + data[i].Id + '">\n' +
                // '                                        <label class="custom-control-label" for="user' + data[i].Id + '"></label>\n' +
                '                                    </div>\n' +
                '                                </li>'

                $("#userList").append(userNode)
                $("#userList2").append(userNode)

            }

        },
    });
}


//通过姓名查找用户
function findUserByName() {
    //用户名
    var name = $(" input[ name='username' ] ").val()
    if (name == "" || name == null) {
        swal("请输入要查询的用户名", "", "error", {
            button: "继续"
        })
    }
    $.ajax({
        type: "POST",
        url: httpUrl+'/user',
        data: {name: name},
        dataType: 'json',
        success: function (data) {
            if (data.Id == 0) {
                swal("不存在此用户", "请重新输入", "info", {
                    button: "继续"
                })
                return;
            }
            $("#userList").children().remove()
            userNode = '                                <li>\n' +
                '                                    <div class="status online"><img src="/static/dist/img/avatars/avatar-female-3.jpg"\n' +
                '                                                                    alt="avatar"><i data-eva="radio-button-on"></i>\n' +
                '                                    </div>\n' +
                '                                    <div class="content">\n' +
                '                                        <h5>' + data.Name + '</h5>\n' +
                '                                        <span>' + data.Email + '</span>\n' +
                '                                    </div>\n' +
                '                                    <div class="custom-control">\n' +
                '                                        <input type="checkbox" onclick="ToFriendUserById(' + data.Id + ')" class="custom-control-input" id="user' + data.Id + '">\n' +
                '                                        <label class="custom-control-label" for="user' + data.Id + '"></label>\n' +
                '                                    </div>\n' +
                '                                </li>'
            $("#userList").append(userNode)
        },
    });

}


//动态增加好友请求模板
function addHandleFriendHtml(user_id, content) {
    var addHandle = '<li">\n' +
        '                                <p >' + content + ' <button style="color: red"  onclick=handleFriendApp(' + user_id + ',"' + content + '")>处理</button></p>\n' +
        '                            </li>'

    //处理通过的模板
    var HandleSuccess = '                            <li>\n' +
        '                                <div class="round"><i data-eva="person-done"></i></div>\n' +
        '                                <p>Quincy has joined to <strong>Squad Ghouls</strong> group.</p>\n' +
        '                            </li>'

    $("#system_notices").append(addHandle)

}


//处理用户申请好友请求
function handleFriendApp(user_id, content) {
    swal({
        title: content,
        text: "是同意呢还是同意呢",
        icon: "info",
        buttons: {
            button1: {
                text: "拒绝",
                value: false,
            },
            button2: {
                text: "同意",
                value: true,
            }
        },
    }).then(function (value) {   //这里的value就是按钮的value值，只要对应就可以啦
        status = value == true ? 1 : 2
        $.ajax({
            type: "POST",
            url: httpUrl+'/user/handleApp',
            data: {from_id: user_id, status: status},
            dataType: 'json',
            success: function (data) {
                GetApplications()
                swal("处理成功", "", "success", {
                    button: "继续"
                })
            },
        });
    });

}


//聊天记录动态添加到html
//type=1给自己加 2给别人加 3 都加
function createChatHtmlNode(content, time, type) {

    var own = '<li class="chat-child">\n' +
        '\n' +
        '                      </li>';

    var other = '<li class="chat-child">\n' +
        '                                        <img src="/static/dist/img/avatars/avatar-male-3.jpg" alt="avatar">\n' +
        '                                        <div class="content">\n' +
        '                                            <div class="message">\n' +
        '                                                <div class="bubble">\n' +
        '                                                    <p>' + content + '\n' +
        '                                                        </p>\n' +
        '                                                </div>\n' +
        '                                            </div>\n' +
        '                                            <span>' + time + '</span>\n' +
        '                                        </div>\n' +
        '                                    </li>'

    if (type == 1) {
        $("#container2").children("ul").append(own)
        $("#container2").children("ul").append(other)
    } else {
        $("#container2").children("ul").append(other)
        $("#container2").children("ul").append(own)
    }
}


//聊天
function sendMessage(to_id, type = 1) {
    //单聊
    var content = $("#chat-text").val()
    //xss
    content = stripscript(content)
    //用来发送之后直接展示的
    content_emm = replace_em(content)
    if (content == "" || content == null) {
        swal("请输入发送内容", "", "error", {
            button: "继续"
        })
        return
    }
    if (sendId == 0) {
        swal("请选择发送对象", "", "error", {
            button: "继续"
        })
        return
    }

    $("#last-message" + sendId).html(content_emm)

    var info = {'message': content, 'type': type, 'to': parseInt(sendId), 'content_type': 1};
    this.wsSend(info);


    $("#chat-text").val("")

    //添加聊天
    createChatHtmlNode(content_emm, '刚刚', 1)
    //下拉框拉到最底部
    msgScrollTOp()

}

function sendEmm() {
    $('.emotion').qqFace({
        id: 'facebox',
        assign: 'chat-text', //给输入框赋值
        path: '/static/emm/arclist/'    //表情图片存放的路径
    });
    var str = $("#chat-text").val();
    em = replace_em(str);
    $("#chat-text").html(em)
}
