var React        = require('react');
var Router       = require('react-router');
var Route        = Router.Route;
var DefaultRoute = Router.DefaultRoute;
var RouteHandler = Router.RouteHandler;
var Link         = Router.Link;

var Artists     = require('./components/Artists.jsx');
var Artist      = require('./components/Artist.jsx');
var Masters     = require('./components/Masters.jsx');
var Master      = require('./components/Master.jsx');
var Label       = require('./components/Label.jsx');
var Styles      = require('./components/Styles.jsx');
var MenuToggle  = require('./components/MenuToggle.jsx');
var Menu        = require('./components/Menu.jsx');
var MenuActions = require('./actions/MenuActions');

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
                        <RouteHandler/>
                    </div>
                </div>
            </div>
        );
    }
});

var routes = (
    <Route handler={App}>
        <DefaultRoute name="index" handler={Artists}/>
        <Route name="artist"  path="artists/:artist_id" handler={Artist}/>
        <Route name="masters" path="masters" handler={Masters}/>
        <Route name="master"  path="masters/:master_id" handler={Master}/>
        <Route name="styles"  path="styles" handler={Styles}/>
        <Route name="label"   path="labels/:label_id" handler={Label}/>
    </Route>
);

Router.run(routes, function (Handler) {
    MenuActions.close();

    React.render(<Handler/>, document.getElementById('app'));
});
