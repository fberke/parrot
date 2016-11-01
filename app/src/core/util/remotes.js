const apiRoot = 'http://localhost:8080/api';

const Remotes = {
    authPath: () => apiRoot + '/authenticate',
    registerPath: () => apiRoot + '/users',
    projectsPath: () => apiRoot + '/projects',
    projectPath: (id) => apiRoot + '/projects/' + id
}

export default Remotes;