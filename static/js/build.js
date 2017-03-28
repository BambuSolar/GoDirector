var Build = (function () {

    var environmnents = {};

    var getEnvironments = function () {

        showLoader();

        var url = '/api/environments';

        $.ajax({
                "url": url,
                "method": 'GET'
            })
            .done(function(result, textStatus, xhr){

                $('#buildEnvironmentSelect')
                    .empty()

                $.each(result.data, function (index, item) {

                    environmnents[item.Name] = item;

                    $('#buildEnvironmentSelect').append(
                        '<option value="' + item.Name + '">' + item.LongName + '</option>'
                    );

                });

                hideLoader();

            })
            .fail(function( jqXHR, textStatus ){

                hideLoader();

                showErrorMessage('An error occurred', 'Please, try it again');

            });

    };

    var listenerFormSubmit = function () {

        $('#form-build').validate({

            rules: {

                Url: {
                    required: true,
                    url: true
                },
                Environment: {
                    required: true
                }

            },

            errorPlacement: function(error, element) {

                var text = $(error).text();

                var parent = $(element).parent();

                parent.removeClass('has-success').addClass('has-error');

                if(text){

                    parent.find('.help-block').remove();

                    parent.append('<span class="help-block">' + text + '</span>');

                }
            },
            success: function(element) {

                var id = $(element).attr('id').split('-')[0];

                var parent = $('#' + id).parent();

                parent.removeClass('has-error').addClass('has-success');

                parent.find('.help-block').remove();

            },
            submitHandler: function(form) {

                var data = JSON.stringify(_getFormData(form));

                console.log(data);

                var method = "POST";

                var url = "/api/builds";

                showLoader();

                $.ajax({
                        "url": url,
                        "method": method,
                        "data": data,
                        "contentType": "application/json"
                    })
                    .done(function(result, textStatus, xhr){

                        hideLoader();

                        console.log(result);

                        //self._successFormRequest(result, textStatus, xhr);

                    })
                    .fail(function( jqXHR, textStatus ){

                        hideLoader();



                        //self._failFormRequest(jqXHR, textStatus );

                    });
            }

        });

    };
    
    var _getFormData = function (form){
        var unindexed_array = $(form).serializeArray();
        var indexed_array = {};

        $.map(unindexed_array, function(n, i){
            indexed_array[n['name']] = n['value'];
        });

        return indexed_array;
    };

    var listeners = function () {


    };

    var init = function () {

        getEnvironments();

        listenerFormSubmit();

        listeners();

    };

    return {
        init: init
    };


}());


$(function () {
    Build.init();
});

