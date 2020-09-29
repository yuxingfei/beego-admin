<div class="row" style="margin-bottom: 15px;">
    <div class="col-sm-8">
        <input type="text" id="captcha" class=" form-control" name="captcha" placeholder="验证码" maxlength="6">
    </div>

    <div class="col-sm-4" style="padding-left: 0">
        <img style="width: 100%;max-width: 120px;" src="{{.captcha.CaptchaUrl}}" alt="图形验证码" id="captchaImg" height="34" onclick="refreshCaptcha()">
    </div>
    <input type="hidden" name="captchaId" id="captchaId" value="{{.captcha.CaptchaId}}">
</div>

<script>
    //刷新验证码
    function refreshCaptcha() {
         id = $("#captchaId").val();
         $.ajax({
             type: "post",
             url: "/admin/auth/refresh_captcha",
             data: {"captchaId":id,"_xsrf":$('meta[name="_xsrf"]').attr('content')},
             dataType: "json",
             success: function(result) {
                 if(result.isNew){
                     $("#captchaImg").attr('src',result.captcha.CaptchaUrl+'?t='+(new Date()).getTime());
                     $("#captchaId").val(result.captcha.CaptchaId);
                 }else{
                     $("#captchaImg").attr('src',$("#captchaImg").attr("src")+"?t="+(new Date()).getTime());
                 }
             }
         });
    }
</script>