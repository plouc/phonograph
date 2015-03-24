var React      = require('react');
var Router     = require('react-router');
var Link       = Router.Link;
var ArtistForm = require('./ArtistForm.jsx');

var ArtistCreate = React.createClass({
    render() {
        return (
            <div>
                <div className="breadcrumbs">
                    <Link to="index">
                        <i className="fa fa-angle-left"/> artists
                    </Link>
                </div>
                <h2 className="page-title">New Artist</h2>
                <ArtistForm/>
            </div>
        );
    }
});

module.exports = ArtistCreate;