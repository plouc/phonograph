var React        = require('react');
var Router       = require('react-router');
var Route        = Router.Route;
var DefaultRoute = Router.DefaultRoute;
var RouteHandler = Router.RouteHandler;
var Link         = Router.Link;

var Artists = require('./components/Artists.jsx');
var Artist  = require('./components/Artist.jsx');

var App = React.createClass({
    getInitialState: function () {
        return {  };
    },

    render: function () {
        return (
            <div>
                <h1><Link to="index">Vinyl API</Link></h1>
                <div>
                    <RouteHandler/>
                </div>
            </div>
        );
    }
});

var routes = (
    <Route handler={App}>
        <DefaultRoute name="index" handler={Artists}/>
        <Route name="artist" path="artists/:artist_id" handler={Artist}/>
    </Route>
);

Router.run(routes, function (Handler) {
    React.render(<Handler/>, document.getElementById('app'));
});
