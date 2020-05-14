
function userInit() {
    let htmlUserInner = document.getElementById("userlogin")
    let htmlDevInner = document.getElementById("devlogin")
    let userLoginContent ='<main>' 
                +'<div class="container">\n'
                + '<div class="login-form">\n'
                + '<form action="/api/v1/user/login" method="post">\n'
                + '<h1>Login for user</h1>\n'
                + '<div class="input-box">\n'
                + '<input type="text" placeholder="username" name="username">\n'
                + '<input type="password" placeholder="password" name="password">\n'
                + '</div>\n'
                + '<div class="btn-box">\n'
                + '<button type="submit">Login</button>\n'
                + '</div>\n'
                + '</form>\n'
                + '<button onclick="changeUserPassInit()">Change Password</button>\n'
                + '</div>\n'
                + '</div>\n'
                + '<main>'
    htmlUserInner.innerHTML = userLoginContent
    htmlDevInner.innerHTML = null
}
function devInit() {
    let htmlDevInner = document.getElementById("devlogin")
    let htmlUserInner = document.getElementById("userlogin")
    let devLoginContent ='<main>' 
                +'<div class="container">\n'
                + '<div class="login-form">\n'
                + '<form action="/api/v1/dev/login" method="post">\n'
                + '<h1>Login for developer</h1>\n'
                + '<div class="input-box">\n'
                + '<input type="text" placeholder="username" name="username">\n'
                + '<input type="password" placeholder="password" name="password">\n'
                + '</div>\n'
                + '<div class="btn-box">\n'
                + '<button type="submit">Login</button>\n'
                + '</div>\n'
                + '</form>\n'
                + '<button onclick="changeDevPassInit()">Change Password</button>\n'
                + '</div>\n'
                + '</div>\n'
                + '<main>'
    htmlDevInner.innerHTML = devLoginContent
    htmlUserInner.innerHTML = null
}



function changeDevPassInit() {
    let htmlUserInner = document.getElementById("userlogin")
    let htmlDevInner = document.getElementById("devlogin")
    let changeDevPassInitContent ='<main>' 
                +'<div class="container">\n'
                + '<div class="login-form">\n'
                + '<form>\n'
                + '<h1>Change password for dev</h1>\n'
                + '<div class="input-box">\n'
                + '<input type="password" placeholder="oldpass" id="oldpass">\n'
                + '<input type="password" placeholder="newpass" id="newpass">\n'
                + '</div>\n'
                + '<div class="btn-box">\n'
                + '</div>\n'
                + '</form>\n'
                + '<button onclick="updateDevPass()">Change</button>\n'
                + '</div>\n'
                + '</div>\n'
                + '<main>'
    htmlDevInner.innerHTML = changeDevPassInitContent
    htmlUserInner.innerHTML = null
}

function updateDevPass() {
    var oldpass = document.getElementById("oldpass").value
    var newpass = document.getElementById("newpass").value
    $.ajax({
        url: '/api/v1/dev/change/pass',
        type: 'POST',
        contentType: 'application/json',
        data:JSON.stringify({
            "oldpass": oldpass,
            "newpass": newpass
        }),
        success: function(data) {
            alert(data)
            window.location.href = "api/v1/auth/login";
        }
    })
}

function changeUserPassInit() {
    let htmlUserInner = document.getElementById("userlogin")
    let htmlDevInner = document.getElementById("devlogin")
    let changeUserPassInitContent ='<main>' 
                +'<div class="container">\n'
                + '<div class="login-form">\n'
                + '<form>\n'
                + '<h1>Change password for user</h1>\n'
                + '<div class="input-box">\n'
                + '<input type="password" placeholder="oldpass" id="oldpass">\n'
                + '<input type="password" placeholder="newpass" id="newpass">\n'
                + '</div>\n'
                + '<div class="btn-box">\n'
                + '</div>\n'
                + '</form>\n'
                + '<button onclick="updateUserPass()">Change</button>\n'
                + '</div>\n'
                + '</div>\n'
                + '<main>'
    htmlUserInner.innerHTML = changeUserPassInitContent
    htmlDevInner.innerHTML = null
}

function updateUserPass() {
    var oldpass = document.getElementById("oldpass").value
    var newpass = document.getElementById("newpass").value
    $.ajax({
        url: '/api/v1/user/change/pass',
        type: 'POST',
        contentType: 'application/json',
        data:JSON.stringify({
            "oldpass": oldpass,
            "newpass": newpass
        }),
        success: function(data) {
            alert(data)
            window.location.href = "api/v1/auth/login";
        }
    })
}


