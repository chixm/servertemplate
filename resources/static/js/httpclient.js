
	// POST でAPIと通信する
	function post(requestUrl, sendData, doAfterSuccess, doAfterFail) {
		$.ajax({
			'url' : requestUrl,
			'dataType':"json",
			'processData': false,
			'data' : sendData,
			'contentType': "application/json;charset=UTF-8",
			'type' : "post",
			'success' : function(result) {
				if (result) {
					doAfterSuccess(result);
				}
			},
			'error' : function(err) {
				console.log(err);
				if (doAfterFail && err.responseJSON) {
					doAfterFail(err.responseJSON);
				}else {
					alert("エラーが発生しました。詳細はconsole logを確認してください。");
				}
			}
		});
	}