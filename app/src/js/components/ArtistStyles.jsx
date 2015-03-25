var React = require('react');

var ArtistStyles = React.createClass({
    render() {
        if (this.props.styles.length === 0) {
            return null;
        }

        var styleNodes = [];
        this.props.styles.forEach((style, i) => {
            styleNodes.push(<span className="artists__list__style" key={style.id}>{style.name}</span>);
            if (i < this.props.styles.length - 1) {
                styleNodes.push(<span>,</span>);
                styleNodes.push(<span>&nbsp;</span>);
            }
        });

        return (
            <div className="artists__list__styles">
                styles: {styleNodes}
            </div>
        );
    }
});

module.exports = ArtistStyles;