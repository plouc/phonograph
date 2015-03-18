var React  = require('react');
var Router = require('react-router');
var Link   = Router.Link;

var MasterReleases = React.createClass({
    render() {
        var listNode = null;
        if (this.props.releases.length > 0) {
            var releaseNodes = this.props.releases.map(release => {
                var labelNodes = release.labels.map(label => {
                    return (
                        <Link className="master__releases__release__label"
                              to="label" params={{ label_id: label.id }}
                        >
                            {label.name}
                        </Link>
                    );
                });

                return (
                    <tr>
                        <td>{release.year}</td>
                        <td>{release.country}</td>
                        <td>{labelNodes}</td>
                    </tr>
                );
            });

            listNode = (
                <table className="master__releases__table">
                    <tbody>
                        {releaseNodes}
                    </tbody>
                </table>
            );
        } else {
            listNode = <p>No release found.</p>
        }

        return (
            <div className="master__releases">
                <h4 className="master__releases__title">Releases</h4>
                {listNode}
            </div>
        );
    }
});

module.exports = MasterReleases;