var React     = require('react');
var Link      = require('react-router').Link;
var MasterRow = require('./MasterRow.jsx');

var ArtistMasters = React.createClass({
    render() {
        var contentNode;
        if (this.props.masters.length > 0) {
            var masterNodes = this.props.masters.map(master => {
                return <MasterRow key={master.id} master={master}/>;
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