$(function(){

    var system_parameters = new CRUD({
        path: '/api/system_parameters',
        listTable: '#index-table',
        nameEntity: 'System Parameter',
        rowTemplateToIndexTable: {
            columns: ['Key', 'Value']
        },
        formFields: [
            {
                field: "Key",
                options: {
                    form: {
                        type: "input:text",
                        label: "Key",
                        placeholder: "Key",
                        autocomplete: false
                    },
                    rule: {
                        required: true,
                        minlength: 2
                    }
                }
            },
            {
                field: "Value",
                options: {
                    form: {
                        type: "input:text",
                        label: "Value",
                        placeholder: "Value",
                        autocomplete: false
                    }
                }
            }
        ]
    });

    system_parameters.init()
});