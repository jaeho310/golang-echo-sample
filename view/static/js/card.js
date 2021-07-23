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