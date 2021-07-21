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
            if (data.success) {
                let html = "";
                let userBox = $("#userBox")
                for (let i = 0; i < data.result.length; i++) {
                    html += '<tr>';
                    html += '<td>' + data.result[i].id + '</td>';
                    html += '<td><a id="detail_btn'+data.result[i].id+'" href="javascript:goToDetail('+data.result[i].id+')">'
                        + data.result[i].name + '</a></td>';
                    html += '</tr>';
                    userList[data.result[i].id] = data.result[i];

                    let option = document.createElement('option');
                    option.innerText = data.result[i].ID;
                    userBox.append(option);
                }
                $("#tableBody").empty();
                $("#tableBody").append(html);
            } else {
                alert(" message = " + data.error)
            }
        },
        error: function(request,status,error){
            alert("code:"+request.status+"\n"+"message:"+request.responseText+"\n"+"error:"+error);
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
            alert("code:"+request.status+"\n"+"message:"+request.responseText+"\n"+"error:"+error);
        }
    })
})