/**
 * 
 * @param {type} url POST请求url
 * @param {type} data 发送POST请求参数
 * @param {type} notice 发请求前显示的确认消息，为空时表示不显示提示
 * @returns {undefined}
 */
function ajaxDev (url,data,notice,closeModel) {
    var result = false;
    if( !notice || confirm(notice)){
        $.post({
          url: url,
          data: data,
          dataType: 'json'
        })
        .done(function (data) {
          if (data.code == 200) {
            //alert(data.msg);
            if(closeModel){
                $('.modal').modal('hide');
            }else{
                window.location.reload();
            }
            result = true;
          } else {
            alert(data.msg);
          }
        });
    }
    return result;
}
