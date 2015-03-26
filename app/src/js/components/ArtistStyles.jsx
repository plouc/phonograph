var React = require('react');
var Link  = require('react-router').Link;

var ArtistStyles = React.createClass({
    render() {
        if (this.props.styles.length === 0) {
            return null;
        }

        var styleNodes = [];
        this.props.styles.forEach((style, i) => {
            styleNodes.push(<Link to="style" params={{ style_id: style.id }} className="artists__list__style" key={style.id}>{style.name}</Link>);
            if (i < this.props.styles.length - 1) {
                styleNodes.push(<span>,</span>);
                styleNodes.push(<span>&nbsp;</span>);
            }
        });

        var classes = 'artist__styles';
        if (this.props.mode === 'list') {
            classes += ' artist__styles--list';
        }

        return (
            <div className={classes}>
                <span className="artist__styles__title">styles</span>
                {styleNodes}
            </div>
        );
    }
});

module.exports = ArtistStyles;