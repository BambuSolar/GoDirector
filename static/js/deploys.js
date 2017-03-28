var Deploy = (function () {

    var environmnents = {};

    var getEnvironments = function () {

        showLoader();

        var url = '/api/environments';

        $.ajax({
                "url": url,
                "method": 'GET'
            })
            .done(function(result, textStatus, xhr){

                $('#buildEnvironmentSelect').empty();

                $.each(result.data, function (index, item) {

                    environmnents[item.Name] = item;

                    $('#environmentSelect').append(
                        '<option value="' + item.Name + '">' + item.LongName + '</option>'
                    );

                    if(item.AllowDirectDeploy) {

                        $('#buildEnvironmentSelect').append(
                            '<option value="' + item.Name + '">' + item.LongName + '</option>'
                        );

                    }

                });
                
                hideLoader();

            })
            .fail(function( jqXHR, textStatus ){

                hideLoader();

                showErrorMessage('An error occurred', 'Please, try it again');

            });

    };

    var enableFormBuild = function (htmlSelect) {

        var environmentSelected = $(htmlSelect).val();

        if(environmentSelected == 'beta'){

            $('#modalBuildBtn').removeAttr('disabled');

        }else{
            $('#modalBuildBtn').attr('disabled','disabled');
        }

    };

    var loadVersions = function (htmlSelect) {

        var environmentSelected = $(htmlSelect).val();

        if(environmentSelected < 0){

            $('#versionSelect')
                .empty()
                .append(
                    '<option value="-1"> -- First select an environment -- </option>'
                );

        }else{

            var url = '/api/builds';

            showLoader();

            $.ajax({
                    "url": url,
                    "method": 'GET'
                })
                .done(function(result, textStatus, xhr){
                    
                    $('#versionSelect').empty();

                    $.each(result.data[environmentSelected], function (index, item) {
                        $('#versionSelect').append(
                            '<option value="' + item + '">' + item + '</option>'
                        );
                    });

                    hideLoader();

                })
                .fail(function( jqXHR, textStatus ){

                    hideLoader();

                    showErrorMessage('An error occurred', 'Please, try it again');

                });


        }

    };

    var listeners = function () {

        $('#environmentSelect').on('change', function(e){
            loadVersions(this);
            enableFormBuild(this);
        });

    };

    var init = function () {

        getEnvironments();

        listeners();

    };

    var anotherMethod = function () {
        // public
    };

    return {
        init: init,
        anotherMethod: anotherMethod
    };


}());


$(function () {
    Deploy.init();
});

