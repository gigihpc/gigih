var App = React.createClass({
  componentWillMount: function() {
    this.setupAjax();
    this.parseHash();
    this.setState();
  },
  // Add access_token if available with each XHR request to API
  setupAjax: function() {
    $.ajaxSetup({
      'beforeSend': function(xhr) {
        if (localStorage.getItem('access_token')) {
          xhr.setRequestHeader('Authorization',
                'Bearer ' + localStorage.getItem('access_token'));
        }
      }
    });
  },
  // Extract the access_token and id_token from Auth0 Callback after login
  parseHash: function(){
    this.auth0 = new auth0.WebAuth({
      domain:       AUTH0_DOMAIN,
      clientID:     AUTH0_CLIENT_ID
    });
    this.auth0.parseHash(window.location.hash, function(err, authResult) {
      if (err) {
        return console.log(err);
      }
      console.log(authResult);
      if(authResult !== null && authResult.accessToken !== null && authResult.idToken !== null){
        localStorage.setItem('access_token', authResult.accessToken);
        localStorage.setItem('id_token', authResult.idToken);
        localStorage.setItem('profile', JSON.stringify(authResult.idTokenPayload));
      }
    });
  },
  // Set user login state
  setState: function(){
    var idToken = localStorage.getItem('id_token');
    if(idToken){
      this.loggedIn = true;
    } else {
      this.loggedIn = false;
    }
  },
  render: function() {

    if (this.loggedIn) {
      return (<LoggedIn />);
    } else {
      return (<Home />);
    }
  }
});