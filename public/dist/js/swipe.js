/**
* Scroll to bottom
*/
window.onload=function () {
     var objDiv = document.getElementById("scroll");
     objDiv.scrollTop = objDiv.scrollHeight;
}
/**
* Switch mode
*/
$(document).ready(function() {
	clicked = true;
	$(".mode").click(function() {
		if(clicked){
			$('head').append('<link href="/static/dist/css/dark.min.css" id="dark" type="text/css" rel="stylesheet">');
			clicked  = false;
		}
		else {
			$('#dark').remove();
			clicked  = true;
		}
	});
});
/**
* Toggle Chat
*/
$(function () {
	'use strict'
	
	$('[data-chat="open"]').on('click', function () {
		$('.chat').toggleClass('open')
	})
})
/**
* Toggle Utility
*/
$(function () {
	'use strict'
	
	$('[data-utility="open"]').on('click', function () {
		$('.utility').toggleClass('open')
	})
})
/**
* Filter
*/
$(document).ready(function(){
	var valueOnLoad = "direct";
	$(".filter").not('.'+valueOnLoad).hide('3000');
	$(".filter").not('.'+valueOnLoad).hide('3000');
	$(".filter-btn").click(function(){
		var value = $(this).attr('data-filter');
		$(".filter").not('.'+value).hide('3000')
		$('.filter').filter('.'+value).show('3000')
	});
});
/**
* Eva Icons
*/
eva.replace()