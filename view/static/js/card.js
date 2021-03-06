$(document).ready(function(){
    console.log("card ready");
    getUsers();
});

let userList = new Object()

function getUsers() {
    $.ajax({
        url: '/api/users',
        contentType: 'application/json',
        type: 'get',
        beforeSend: function(xhr) {
            // xhr.setRequestHeader(header, token);
        },
        success: function(data) {
            if (data.success) {
                let userBox = $("#userBox")
                for (let i = 0; i < data.result.length; i++) {
                    let option = document.createElement('option');
                    option.innerText = data.result[i].name;
                    userBox.append(option);
                    userList[data.result[i].name] = data.result[i];
                }
            } else {
                alert(" message = " + data.error)
            }
        },
        error: function(request,status,error){
            alert("code:"+request.status+"\n"+"message:"+request.responseText+"\n"+"error:"+error);
        }
    })
}

$("#cancelBtn").on("click", function() {
    window.location.href="/list"
});

$("#regBtn").on("click", function() {
    let cardName = $("#cardName").val();
    let cardLimit = $("#cardLimit").val();
    let userName = $("#userBox").val()
    isValidate = validationCheck(cardName, cardLimit)
    if (!isValidate) {
        return
    }
    let data = new Object()
    data["name"] = cardName
    data["limit"] = Number(cardLimit)
    data["userId"] = userList[userName].id
    $.ajax({
        url: '/api/cards',
        contentType: 'application/json',
        type: 'post',
        data: JSON.stringify(data),
        beforeSend: function(xhr) {
            // xhr.setRequestHeader(header, token);
        },
        success: function(data) {
            alert("등록성공")
            window.location.href="/list"
        },
        error: function(request,status,error){
            alert("code:"+request.status+"\n"+"message:"+request.responseText+"\n"+"error:"+error);
        }
    })
});

function validationCheck(cardName, cardLimit) {
    if (!cardName || !cardLimit) {
        alert("빈칸을 모두 입력해주세요")
        return false
    }
    if (isNaN(cardLimit)) {
        alert("한도에는 숫자만 입력해주세요")
        return false
    }
    return true
}