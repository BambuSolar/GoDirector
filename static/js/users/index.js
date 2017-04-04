$(function(){

    var system_parameters = new CRUD({
        path: '/api/users',
        listTable: '#index-table',
        nameEntity: 'Users',
        operations: ['delete'],
        rowTemplateToIndexTable: {
            columns: ['FullName', 'Email']
        }
    });

    system_parameters.init()
});