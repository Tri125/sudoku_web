$(document).ready(function() {

	//Select the text input when clicking it
	$("input[type='text']").click(function(){
		$(this).select();
	});

	//Select the underlying text input when clicking on the td element
	$('td').click(function(){
		$(this).children("input[type='text']").select();
	});

	//Prevent the event to be handled when an invalid key is pressed
	$("input[type='text']").keypress(function(event){
		if (event.which < 49 || event.which > 57)
		{
			event.preventDefault();
		}
	});

	//Remove focus if pressed a valid key ranging from 1-9
	$("input[type='text']").keyup(function(event){
		if (event.which >= 49 && event.which <= 57)
		{
			$(this).blur();
			//Select next cell
			moveNext(this);
		}
	});

	//When pressing enter, move to the next cell
	$("input[type='text'").keydown(function(event){
		if (event.which == 13)
		{
			moveNext(this);
		}
	});

	function moveNext(element)
	{
		var trNext = $(element).parent().parent('tr').next();

		var nextElement = $(element).parent().next();

		var nextTBody = $(element).parent().parent().parent('tbody').next();

		if (nextElement.length !== 0)
		{
			$(nextElement).children().first().select();
		}
		else if (trNext.length !== 0)
		{
 			$(trNext).children().first().children().first().select();
		}
		else if (nextTBody.length !== 0)
		{
			$(nextTBody).children().first().children().first().children().first().select();
		}
	}
});