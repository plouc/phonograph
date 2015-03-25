var React        = require('react');
var Router       = require('react-router');
var Route        = Router.Route;
var DefaultRoute = Router.DefaultRoute;
var RouteHandler = Router.RouteHandler;
var Link         = Router.Link;

var Artists      = require('./components/Artists.jsx');
var ArtistCreate = require('./components/ArtistCreate.jsx');
var Artist       = require('./components/Artist.jsx');
var Masters      = require('./components/Masters.jsx');
var Master       = require('./components/Master.jsx');
var Label        = require('./components/Label.jsx');
var Styles       = require('./components/Styles.jsx');
var Skills       = require('./components/Skills.jsx');
var Skill        = require('./components/Skill.jsx');
var MenuToggle   = require('./components/MenuToggle.jsx');
var Menu         = require('./components/Menu.jsx');
var MenuActions  = require('./actions/MenuActions');
var Promise      = require('bluebird');

var App = React.createClass({
    render: function () {
        return (
            <div>
                <div className="header">
                    <Link className="header__brand" to="index"><i className="fa fa-dot-circle-o"/> phonograph</Link>
                    <MenuToggle/>
                </div>
                <Menu />
                <div>
                    <div className="container">
                        <RouteHandler {...this.props}/>
                    </div>
                </div>
            </div>
        );
    }
});

var routes = (
    <Route handler={App}>
        <DefaultRoute name="index" handler={Artists}/>

        <Route name="artist_create" path="artists/new"        handler={ArtistCreate}/>
        <Route name="artist"        path="artists/:artist_id" handler={Artist}/>
        <Route name="masters"       path="masters"            handler={Masters}/>
        <Route name="master"        path="masters/:master_id" handler={Master}/>
        <Route name="styles"        path="styles"             handler={Styles}/>
        <Route name="skills"        path="skills"             handler={Skills}/>
        <Route name="skill"         path="skills/:skill_id"   handler={Skill}/>
        <Route name="label"         path="labels/:label_id"   handler={Label}/>
    </Route>
);

function fetchData(routes, state) {
    var data = {};
    var promises = routes
        .filter(route => route.handler.fetchData)
        .map(route => {
            return route.handler.fetchData(state.params, state.query).then(d => {
                data[route.name] = d;
            });
        })
    ;

    return Promise.all(promises).then(() => data);
}

Router.run(routes, function (Handler, state) {
    fetchData(state.routes, state)
        .then(data => {
            React.render(<Handler data={data}/>, document.getElementById('app'));
        })
        .catch(err => {
            console.log('ERR', err);
        })
    ;
});
