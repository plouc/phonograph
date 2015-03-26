var React = require('react');
var Link  = require('react-router').Link;

var MasterRow = React.createClass({
    render() {
        var master = this.props.master;

        var imgNode = null;
        if (master.picture !== '') {
            imgNode = (
                <div className="masters__list__item__picture">
                    <img src={`/images/${ master.picture }`}/>
                </div>
            );
        }

        return (
            <div className="masters__list__item">
                {imgNode}
                <Link className="masters__list__item__name" to="master" params={{ master_id: master.id }}>{master.name}</Link>
                <span>{master.year}</span>
            </div>
        );
    }
});

module.exports = MasterRow;