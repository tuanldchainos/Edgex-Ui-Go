$(document).ready(function () {

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
			url: '/api/v1/user/logout',
			type: 'GET',
			success: function () {
				window.location.href = '/api/v1/auth/login'
			}
		});
	});

	$.ajax({
		url: "/data/user-menu.json",
		type: "GET",
		dataType: "json",
		success: function (data) {
			menuRender(data);
		}
	})

	function menuRender(data) {
		for(var i=0; i < data.length; i++) {
			var menu = data[i];
            var str
            if(menu.status==true){
               str = '<li class="nav-item">' + '<a class="nav-link active" id="' + menu.title +'-tab" data-toggle="tab" href="#' +menu.title+ '" role="tab" aria-controls="'+ menu.title +'" aria-selected="'+ menu.status+'">' + menu.title + '</li>'
            }else {
                str = '<li class="nav-item">' + '<a class="nav-link" id="' + menu.title +'-tab" data-toggle="tab" href="#' +menu.title+ '" role="tab" aria-controls="'+ menu.title +'" aria-selected="'+ menu.status+'">' + menu.title + '</li>'
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
        $('#content #' + title).load(url)
	}
	
	window.addEventListener("unload", function(){
		$.ajax({
			url: '/api/v1/user/logout',
			type: 'GET',
			success: function () {
				window.location.href = '/api/v1/auth/login'
			}
		})
	})
})
