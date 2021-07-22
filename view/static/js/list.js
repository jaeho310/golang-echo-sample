// let token = $("meta[name='_csrf']").attr("content");
// let header = $("meta[name='_csrf_header']").attr("content");
let userList = new Object();

$(document).ready(function(){
    console.log("list ready");
    getUser();
});

function getUser() {
    $.ajax({
        url: '/api/users',
        contentType: 'application/json',
        type: 'get',
        beforeSend: function(xhr) {
            // xhr.setRequestHeader(header, token);
        },
        success: function(data) {
            let html = "";
            let userBox = $("#userBox")
            for (let i = 0; i < data.length; i++) {
                html += '<tr>';
                html += '<td>' + data[i].id + '</td>';
                html += '<td><a id="detail_btn'+data[i].id+'" href="javascript:goToDetail('+data[i].id+')">' + data[i].name + '</a></td>';
                html += '</tr>';
                userList[data[i].id] = data[i];

                let option = document.createElement('option');
                option.innerText = data[i].id;
                userBox.append(option);
            }
            $("#tableBody").empty();
            $("#tableBody").append(html);
        },
        error: function(request,status,error){
            alert(" message = " + request.responseText);
        }
    })
}

function goToDetail(n) {
    let url = "/detail/" + n
    window.location.href = url
}

$("#deleteBtn").on("click",function() {
    let result = confirm("삭제하시겠습니까?");
    if (result == false) {
        return;
    }
    $.ajax({
        url: '/api/users/' + $("#userBox").val(),
        contentType: 'application/json',
        type: 'delete',
        success: function(data) {
            alert("삭제되었습니다.");
            location.reload();
        },
        error: function(request,status,error){
            alert(" message = " + request.responseText);
        }
    })
})