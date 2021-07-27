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

function reloadCardBox() {
    let cardBox = $("#cardBox")
    cardBox.find("option").remove();
    let option = document.createElement('option');
    option.innerText = "삭제할 카드를 선택해주세요"
    cardBox.append(option);
}

$("#userBox").change(function (){
    reloadCardBox();
    let userName = $("#userBox").val();
    if (userName == "사용자를 선택해주세요") {
        alert("사용자를 선택해주세요")
        return;
    }
    fillCardBox(userName);
})

function fillCardBox(userName) {
    let cardBox = $("#cardBox")
    let cards = userList[userName].cards;
    for (let i = 0; i < cards.length; i++) {
        let option = document.createElement('option');
        option.innerText = cards[i].name;
        cardBox.append(option);
    }
}

$("#cancelBtn").on("click", function() {
    window.location.href="/list"
});

function getCardId(cardName) {
    let cards = userList[$("#userBox").val()].cards;
    for (let i = 0; i < cards.length; i++) {
        if (cardName == cards[i].name) {
            return cards[i].id
        }
    }

}

$("#deleteBtn").on("click", function() {
    let userId = userList[$("#userBox").val()].id;
    let cardName = $("#cardBox").val();
    if (cardName == "삭제할 카드를 선택해주세요") {
        alert("삭제할 카드를 선택해주세요")
        return;
    }
    let cardId = getCardId(cardName);
    $.ajax({
        url: '/api/cards?cardId=' + cardId + "&userId=" +userId,
        contentType: 'application/json',
        type: 'delete',
        beforeSend: function(xhr) {
            // xhr.setRequestHeader(header, token);
        },
        success: function(data) {
            alert("삭제성공")
            window.location.href="/list"
        },
        error: function(request,status,error){
            alert("code:"+request.status+"\n"+"message:"+request.responseText+"\n"+"error:"+error);
        }
    })
});