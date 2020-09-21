<div class="row" style="margin-bottom: 15px;">
    <div class="col-sm-8">
        <input type="text" id="captcha" class=" form-control" name="captcha" placeholder="验证码" maxlength="6">
    </div>

    <div class="col-sm-4" style="padding-left: 0">
        <img style="width: 100%;max-width: 120px;" src="/admin/auth/captcha/{{.captcha_id}}.png" alt="图形验证码" id="captchaImg" height="34" onclick="refreshCaptcha('captchaImg')">
    </div>
    <input type="hidden" name="captchaId" value="{{.captcha_id}}">
</div>

<script>

    function refreshCaptcha(dom) {
        let $dom = $('#'+dom);
        $dom.attr('src','/admin/auth/captcha/{{.captcha_id}}.png?reload='+(new Date()).getTime());
    }
</script>