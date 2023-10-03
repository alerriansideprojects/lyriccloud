class App extends React.Component {
  render() {
    return <Home />;
  }
}

class Home extends React.Component {
  render() {
    return (
      <div className="container">
        <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
          <h1>Home</h1>
          <p>This is the home page</p>
        </div>
      </div>
    );
  }
}

class Artists extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      artists: [],
    };
  }

  render() {
    return (
      <div className="container">
        <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
          <h1>Artists</h1>
          <p>A list of artists</p>
          <div className="row">
            {this.state.artists.map((artist) => {
              return (
                <div className="col-xs-4" key={artist.id}>
                  <a href={`/#/artists/${artist.id}`}>
                    <h4>{artist.name}</h4>
                  </a>
                </div>
              );
            })}
          </div>
        </div>
      </div>
    );
  }
}

ReactDOM.render(<App />, document.getElementById("app"));
