var Home = React.createClass({
  // Clicking the login link will redirect the user to the Hosted Lock page 
  // where the user will provide their credentials and will then be redirected
  // back to the application
  authenticate: function(){
    this.webAuth = new auth0.WebAuth({
      domain:       AUTH0_DOMAIN,
      clientID:     AUTH0_CLIENT_ID,
      scope:        'openid profile',
      audience:     AUTH0_API_AUDIENCE,
      responseType: 'token id_token',
      redirectUri : AUTH0_CALLBACK_URL
    });
    this.webAuth.authorize();
  },
  render: function() {
    return (
    <div className="container">
      <div className="col-xs-12 jumbotron text-center">
        <h1>We R VR</h1>
        <p>Provide valuable feedback to VR experience developers.</p>
        <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
      </div>
    </div>);
  }
});