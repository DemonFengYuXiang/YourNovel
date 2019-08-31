/**
 * Created by howie on 21/02/2017.
 */
$(document).ready(function () {
    var link_prefix = $("#link_prefix").val();
    var chapter_url = $("#url").val();
    var novels_name = $("#novels_name").val();
    var domain = $("#domain").val();
    $(".container a").each(function () {
        var url = $(this).attr('href');
        url = url.trimLeft("/");
        if (typeof(url) != "undefined") {
            if (url.indexOf("yournovel") < 0) {
                var name = $(this).text();
                var prefix = '';
                if (link_prefix == '-1') {
                    prefix = chapter_url;
                }else if (link_prefix == '1') {
                    prefix = '';
                }else if (link_prefix == '0') {
                   prefix = domain;
                }
                show_url = "content?content_url=" + prefix + '/' + url + "&content_title=" + name + "&chapter_url=" + chapter_url + "&novel_name=" + novels_name;
                $(this).attr('href', show_url);
                $(this).attr('target', '_blank');
            }
        }
    });
});