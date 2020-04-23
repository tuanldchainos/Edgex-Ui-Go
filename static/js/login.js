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
                + '<button>Change Password</button>\n'
                + '</div>\n'
                + '</form>\n'
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
                + '</div>\n'
                + '</div>\n'
                + '<main>'
    htmlDevInner.innerHTML = devLoginContent
    htmlUserInner.innerHTML = null
}