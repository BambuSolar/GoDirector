$(function(){

    var environments = new CRUD({
        path: '/api/environments',
        listTable: '#index-table',
        rowTemplateToIndexTable: {
            columns: ['Name', 'Version', 'Branch']
        }
    });

    environments.init()
});