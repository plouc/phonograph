var React = require('react');
var Link  = require('react-router').Link;

var StyleRow = React.createClass({
    render() {
        return (
            <div className="list__item">
                <Link to="style" params={{ style_id: this.props.style.id }} key={this.props.style.id}>{this.props.style.name}</Link>
            </div>
        );
    }
});

module.exports = StyleRow;