$(document).ready(function () {
	debugger
	$.ajaxSetup({
		cache: false,
		headers: { "X-Session-Token": window.sessionStorage.getItem("X_Session_Token") },
		statusCode: {
			302: function () {
				window.location.href = '/api/v1/auth/login'; //prevent browser cache result to redirect  failed.
			}
		}
	});
    
    //logout control
	$(".headbar li.logout").on("click", function () {
		$.ajax({
			url: '/api/v1/dev/logout',
			type: 'GET',
			success: function () {
				window.location.href = '/api/v1/auth/login'
			}
		});
	});

	$.ajax({
		url: "/data/dev-menu.json",
		type: "GET",
		dataType: "json",
		success: function (data) {
			menuRender(data);
		}
	})

	function menuRender(data) {
		debugger
		for(var i = 0; i < data.length; i++) {
			var menu = data[i];
			console.log(menu)
            var str
            if(menu.status==true){
               str = '<li class="nav-item">' + '<a href="#'+ menu.title + '" class="nav-link active" id="' + menu.title +'-tab" data-toggle="tab" role="tab" aria-controls="'+ menu.title +'" aria-selected="'+ menu.status+'">' + menu.title + '</li>'
            }else {
				str = '<li class="nav-item">' + '<a href="#'+ menu.title + '" class="nav-link" id="' + menu.title +'-tab" data-toggle="tab" role="tab" aria-controls="'+ menu.title +'" aria-selected="'+ menu.status+'">' + menu.title + '</li>'
            }
			$('#menutab').append(str)
			createTabByTitle(menu.title, menu.url, menu.status)
		}
	}


	function createTabByTitle(title, url, status) {
        var str
        if (status == true) {
            str = '<div class="tab-pane fade show active" id="' + title + '" role="tabpanel" aria-labelledby="' + title + '-tab"></div>'
        }else {
            str = '<div class="tab-pane fade" id="' + title + '" role="tabpanel" aria-labelledby="' + title + '-tab"></div>'
        }
		$('#content').append(str)
		console.log(url)
		$('#content #' + title).load(url, null, null)
	}

	window.addEventListener("unload", function(){
		$.ajax({
			url: '/api/v1/dev/logout',
			type: 'GET',
			success: function () {
				window.location.href = '/api/v1/auth/login'
			}
		})
	})
})
