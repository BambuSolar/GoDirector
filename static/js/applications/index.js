$(function(){

    var system_parameters = new CRUD({
        path: '/api/applications',
        listTable: '#index-table',
        nameEntity: 'Applications',
        operations: ['show', 'edit', 'delete','new'],
        rowTemplateToIndexTable: {
            columns: ['Name', 'IP']
        },
        formFields: [
            {
                field: "Name",
                options: {
                    form: {
                        type: "input:text",
                        label: "Application Name",
                        placeholder: "Application Name",
                        autocomplete: false
                    },
                    rule: {
                        required: true,
                        minlength: 3
                    }
                }
            },
            {
                field: "IP",
                options: {
                    form: {
                        type: "input:text",
                        label: "IP Address",
                        placeholder: "IP Address",
                        autocomplete: false
                    },
                    rule: {
                        required: true
                    }
                }
            }
        ]
    });

    system_parameters.init()
});