package services

import (
	"beego-admin/utils"
	"strings"
)

// SettingFormService struct
type SettingFormService struct {
	BaseService
}

// formHTML 生成form html
var formHTML map[string]string = map[string]string{
	"checkboxHtml": `<div class="form-group">
		<label class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<div class="checkbox">
				<label>
					<input type="checkbox" name="[FIELD_NAME][]" class="field-checkbox"> [FORM_NAME]
				</label>
			</div>
		</div>
	</div>`,

	"colorHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<div class="input-group" id="color-[FIELD_NAME]">
				<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请输入[FORM_NAME]" type="text" class="form-control field-map">
				<div class="input-group-addon"><i></i></div>
			</div>
		</div>
	</div>
	<script>
		$('#color-[FIELD_NAME]').colorpicker();
	</script>`,

	"dateHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请选择[FORM_NAME]" type="text" class="form-control filed-date">
		</div>
	</div>
	<script>
		laydate.render({
			elem: '#[FIELD_NAME]',
		});
	</script>`,

	"dateRangeHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请选择[FORM_NAME]" type="text" class="form-control filed-date-range">
		</div>
	</div>
	<script>
		laydate.render({
			elem: '#[FIELD_NAME]',
			range: true
		});
	</script>`,

	"datatimeHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请选择[FORM_NAME]" type="text" class="form-control filed-datetime">
		</div>
	</div>
	<script>
		laydate.render({
			elem: '#[FIELD_NAME]',
			type: 'datetime',
		});
	</script>`,

	"datetimeRangeHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请选择[FORM_NAME]" type="text" class="form-control filed-datetime-range">
		</div>
	</div>
	<script>
		laydate.render({
			elem: '#[FIELD_NAME]',
			type: 'datetime',
			range: true,
		});
	</script>`,

	"EditorHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
			<div class="col-sm-10">
				<script id="[FIELD_NAME]" name="[FIELD_NAME]" type="text/plain">{\$data.[FIELD_NAME]|raw|default='[FIELD_DEFAULT]'}</script>
			</div>
		</div>
	<script>
		UE.delEditor('[FIELD_NAME]');
		var UE_[FIELD_NAME] = UE.getEditor('[FIELD_NAME]',{
			serverUrl :UEServer
		});
	</script>`,

	"emailhtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请输入[FORM_NAME]" type="email" class="form-control field-email">
		</div>
	</div>`,

	"fileHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4"> 
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" data-initial-preview="{\$data.[FIELD_NAME]|default=''}" placeholder="请上传[FORM_NAME]" type="file" class="form-control field-file" >
		</div>
	</div>
	<script>
		$('#[FIELD_NAME]').fileinput({
			language: 'zh',
			browseLabel: '浏览',
			initialPreviewAsData: false,
			dropZoneEnabled: false,
			showUpload:false,
			showRemove: false,
			allowedFileExtensions: ['jpg', 'png', 'gif','bmp','svg','jpeg','mp4','doc','docx','pdf','xls','xlsx','ppt','pptx','txt'],
			maxFileSize:10240
		});
	</script>`,

	"iconHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<div class="input-group iconpicker-container">
				<span class="input-group-addon"><i class="fa fa-pencil"></i></span>
				<input maxlength="30" id="[FIELD_NAME]" name="[FIELD_NAME]"
					   value="[FIELD_CONTENT]" class="form-control "
					   placeholder="请选择[FORM_NAME]">
			</div>
		</div>
	</div>
	<script>
		$('#[FIELD_NAME]').iconpicker({placement: 'bottomLeft'});
	</script>`,

	"idCardHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请输入[FORM_NAME]" type="text" maxlength="18" class="form-control field-id-card">
		</div>
	</div>`,

	"imageHtml": `<div class="form-group">
        <label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
        <div class="col-sm-10 col-md-4"> 
            <input id="[FIELD_NAME]" name="[FIELD_NAME]"  placeholder="请上传[FORM_NAME]" data-initial-preview="[FIELD_CONTENT]" type="file" class="form-control field-image" >
        </div>
    </div>
    <script>
    $('#[FIELD_NAME]').fileinput({
        language: 'zh',
        overwriteInitial: true,
        browseLabel: '浏览',
        initialPreviewAsData: true,
        dropZoneEnabled: false,
        showUpload:false,
        showRemove: false,
        allowedFileTypes:['image'],
        maxFileSize:10240,
    });
    </script>`,

	"ipHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请输入[FORM_NAME]" type="text" class="form-control field-map">
		</div>
	</div>`,

	"ampHtml": `<div class="form-group">
		<label class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-8 ">
			<div id="map-container" style="width: 100%; height: 350px;position: relative; background-color: rgb(229, 227, 223);overflow: hidden;transform: translateZ(0px);">
			</div>
			<input name="[FIELD_NAME_LNG]" hidden id="[FIELD_NAME_LNG]" value="{\$data.[FIELD_NAME_LNG]|default='[FIELD_DEFAULT_LNG]'}">
			<input name="[FIELD_NAME_LAT]" hidden id="[FIELD_NAME_LAT]" value="{\$data.[FIELD_NAME_LAT]|default='[FIELD_DEFAULT_LAT]'}" >
		</div>
	</div>
	<script>
		AMapUI.loadUI(['misc/PositionPicker'], function(PositionPicker) {
			var map = new AMap.Map('map-container', {
				zoom: 16,
				scrollWheel: true
			})
			var positionPicker = new PositionPicker({
				mode: 'dragMap',
				map: map
			});
			positionPicker.on('success', function(positionResult) {
				console.log(positionResult);
				console.log('success');
				$('#[FIELD_NAME_LNG]').val(positionResult.position.lng);
				$('#[FIELD_NAME_LAT]').val(positionResult.position.lat);
			});
			positionPicker.on('fail', function(positionResult) {
				console.log(positionResult);
			});
			positionPicker.start( 
				{if isset(\$data)}
				new AMap.LngLat({\$data.[FIELD_NAME_LNG]}, {\$data.[FIELD_NAME_LAT]})
				{/if}
				); 
			map.panBy(0, 1);
			map.addControl(new AMap.ToolBar({
				liteStyle: true
			}))
		});
	</script>`,

	"mobileHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请输入[FORM_NAME]" type="tel" maxlength="11" class="form-control field-mobile">
		</div>
	</div>`,

	"multiFileHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-8"> 
			<input id="[FIELD_NAME]" name="[FIELD_NAME]"  placeholder="请上传[FORM_NAME]" type="file" class="form-control field-multi-file" >
		</div>
	</div>
	<script>
		$('#[FIELD_NAME]').fileinput({
			//theme: 'fas',
			language: 'zh',
		
			browseLabel: '浏览',
			initialPreviewAsData: false,
			initialPreviewShowDelete:false,
			dropZoneEnabled: false,
			showUpload:false,
			showRemove: false,
			allowedFileExtensions: ['jpg', 'png', 'gif','bmp','svg','jpeg','mp4','doc','docx','pdf','xls','xlsx','ppt','pptx','txt'],
			//默认限制10M
			maxFileSize:10240
		});
	</script>`,

	"multiImageHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4"> 
			<input id="[FIELD_NAME]" name="[FIELD_NAME][]"  placeholder="请上传[FORM_NAME]" multiple="multiple" type="file" class="form-control field-multi-image" >
		</div>
	</div>
	<script>
	$(function() {
		$('#[FIELD_NAME]').fileinput({
			"initialPreview":false,
			overwriteInitial: true,
			language: 'zh',
			browseLabel: '浏览',
			initialPreviewAsData: true,
			initialPreviewShowDelete:false,
			dropZoneEnabled: false,
			showUpload:false,
			showRemove: false,
			allowedFileTypes:['image'],
			//默认限制10M
			maxFileSize:10240,
		});
	})
	</script>`,

	"multiSelectHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<select name="[FIELD_NAME][]" id="[FIELD_NAME]" data-placeholder="请选择[FORM_NAME]" class="form-control field-multi-select" multiple="multiple">
				<option value=""></option>
			</select>
		</div>
	</div>
	<script>
	 $('#[FIELD_NAME]').select2();
	</script>`,

	"numberHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			 <div class="input-group">
				<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请输入[FORM_NAME]" type="number" class="form-control field-number">
			</div>
		</div>
	</div>
	<script>
		$('#[FIELD_NAME]')
			.bootstrapNumber({
				upClass: 'success',
				downClass: 'primary',
				center: true
			});
	</script>`,

	"passwordHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请输入[FORM_NAME]" type="password" class="form-control field-password">
		</div>
	</div>`,

	"radioHtml": `<div class="radio">
		<label>
			<input type="radio" name="[FIELD_NAME]" value="" checked="">
			[FORM_NAME]
		</label>
	</div>`,

	"selectHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<select name="[FIELD_NAME]" id="[FIELD_NAME]" class="form-control field-select" data-placeholder="请选择[FORM_NAME]">
				<option value=""></option>
				[OPTION_DATA]
			</select>
		</div>
	</div>
	<script>
	 $('#[FIELD_NAME]').select2();
	</script>`,

	"switchHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
		<input class="input-switch"  id="[FIELD_NAME]" value="1" [SWITCH_CHECKED] type="checkbox" />
		<input class="switch field-switch" placeholder="[FORM_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" hidden />
		</div>
	</div>
	<script>
		$('#[FIELD_NAME]').bootstrapSwitch({
			onText: "[ON_TEXT]",
			offText: "[OFF_TEXT]",
			onColor: "success",
			offColor: "danger",
			onSwitchChange: function (event, state) {
				$(event.target).closest('.bootstrap-switch').next().val(state ? '1' : '0').change();
			}
		});
	</script>`,

	"textHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请输入[FORM_NAME]" type="text" class="form-control field-text">
		</div>
	</div>`,

	"textareaHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<textarea id="[FIELD_NAME]" name="[FIELD_NAME]" class="form-control" rows="[ROWS]" placeholder="请输入[FORM_NAME]">[FIELD_CONTENT]</textarea>
		</div>
	</div>`,

	"timeHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请选择[FORM_NAME]" type="text" class="form-control filed-time">
		</div>
	</div>
	<script>
		laydate.render({
			elem: '#[FIELD_NAME]',
			type: 'time',
		});
	</script>`,

	"timeRangeHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请选择[FORM_NAME]" type="text" class="form-control filed-time-range">
		</div>
	</div>
	<script>
		laydate.render({
			elem: '#[FIELD_NAME]',
			type: 'time',
			range: true,
		});
	</script>`,

	"urlHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请输入[FORM_NAME]" type="text" class="form-control field-map">
		</div>
	</div>`,

	"yearHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请选择[FORM_NAME]" type="text" class="form-control filed-year">
		</div>
	</div>
	<script>
		laydate.render({
			elem: '#[FIELD_NAME]',
			type: 'year',
		});
	</script>`,

	"yearMonthHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请选择[FORM_NAME]" type="text" class="form-control filed-year-month">
		</div>
	</div>
	<script>
		laydate.render({
			elem: '#[FIELD_NAME]',
			type: 'month',
		});
	</script>`,

	"yearMonthRangeHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请选择[FORM_NAME]" type="text" class="form-control filed-year-month-range">
		</div>
	</div>
	<script>
		laydate.render({
			elem: '#[FIELD_NAME]',
			type: 'month',
			range: true,
		});
	</script>`,

	"yearRangeHtml": `<div class="form-group">
		<label for="[FIELD_NAME]" class="col-sm-2 control-label">[FORM_NAME]</label>
		<div class="col-sm-10 col-md-4">
			<input id="[FIELD_NAME]" name="[FIELD_NAME]" value="[FIELD_CONTENT]" placeholder="请选择[FORM_NAME]" type="text" class="form-control filed-year-range">
		</div>
	</div>
	<script>
		laydate.render({
			elem: '#[FIELD_NAME]',
			type: 'year',
			range: true,
		});
	</script>`,
}

// GetFieldForm 获取表单字段
func (*SettingFormService) GetFieldForm(typeS string, name string, field string, content string, option string) string {
	htmlVar := utils.ParseName(typeS, 1, false) + "Html"

	form, ok := formHTML[htmlVar]

	if !ok {
		return ""
	}

	switch typeS {
	case "switch":
		form = strings.ReplaceAll(form, "[ON_TEXT]", "是")
		form = strings.ReplaceAll(form, "[OFF_TEXT]", "否")
		if content == "" || content == "0" {
			form = strings.ReplaceAll(form, "[SWITCH_CHECKED]", "")
		} else {
			form = strings.ReplaceAll(form, "[SWITCH_CHECKED]", "checked")
		}
		break
	case "select":
		optionHTML := ""
		optionArr := strings.Split(option, "\r\n")
		for _, item := range optionArr {
			optionKeyValueArr := strings.Split(item, "||")
			selectStr := ""
			if len(optionKeyValueArr) > 0 && content == optionKeyValueArr[0] {
				selectStr = "selected"
			}
			if len(optionKeyValueArr) > 0 {
				optionHTML += `<option value="` + optionKeyValueArr[0] + `" ` + selectStr + `>` + optionKeyValueArr[1] + `</option>`
			}
		}
		form = strings.ReplaceAll(form, "[OPTION_DATA]", optionHTML)
		break
	default:
		break
	}

	form = strings.ReplaceAll(form, "[FIELD_NAME]", field)
	form = strings.ReplaceAll(form, "[FORM_NAME]", name)
	form = strings.ReplaceAll(form, "[FIELD_CONTENT]", content)
	return form
}
