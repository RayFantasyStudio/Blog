/**
 * Created by Allen on 2016/5/30.
 */

$('.ui.dropdown')
    .dropdown()
;

$('.ui.accordion')
    .accordion()
;
$('.ui.radio.checkbox')
    .checkbox()
;
var tag_value;
tag_value = $('#tag_selector').dropdown('get value');
$('#tag_js_value').value = tag_value;