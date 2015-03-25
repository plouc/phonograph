var React = require('react');
var Link  = require('react-router').Link;

var ArtistMasters = React.createClass({
    render() {
        var contentNode;
        if (this.props.masters.length > 0) {
            var masterNodes = this.props.masters.map(master => {
                return (
                    <Link key={master.id} to="master" params={{ master_id: master.id }} className="list__item">
                        <span className="list__item__label">{master.name}</span>
                        <span className="master__year">{master.year}</span>
                    </Link>
                );
            });
            contentNode = (
                <div className="list">
                    {masterNodes}
                </div>
            );
        } else {
            contentNode = <p>No item found.</p>
        }

        return (
            <div className="artist__masters">
                <h4 className="artist__masters__title">Played in</h4>
                {contentNode}
            </div>
        );
    }
});

module.exports = ArtistMasters;