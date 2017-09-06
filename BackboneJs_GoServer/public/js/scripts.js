Backbone.Model.prototype.idAttribute = '_id';

// Backbone Model

var Mhsw = Backbone.Model.extend({
	default:{
	//	_id: "1",
		name: '',
		address: '',
		old: ''
	},
	//idAttribute:"_id"
});

// Backbone Collection

var Mhsws = Backbone.Collection.extend({
	url: 'http://localhost:3003/api/mhsws'
});

// instantiate two Mhsws

/*var mhsw1 = new Mhsw({
	name: 'M1',
	address: 'tglk1',
	old: '1'
});
mhsw1.save()

var mhsw2 = new Mhsw({
	name: 'M2',
	address: 'tglk2',
	old: '2'
});

mhsw1.save(null, {
			success: function(response) {
				console.log('Successfully SAVED mhsw with _id: ' + response.toJSON()._id);
			},
			error: function() {
				console.log('Failed to save mhsw!');
			}
		});
mhsw2.save(null, {
			success: function(response) {
				console.log('Successfully SAVED mhsw with _id: ' + response.toJSON()._id);
			},
			error: function() {
				console.log('Failed to save mhsw!');
			}
		});*/
// instantiate a Collection

var mhsws = new Mhsws();

// Backbone View for one blog

var MhswView = Backbone.View.extend({
	model: new Mhsw(),
	tagName: 'tr',
	initialize: function() {
		this.template = _.template($('.mhsw-list-template').html());
	},
	events: {
		'click .edit-mhsw': 'edit',
		'click .update-mhsw': 'update',
		'click .cancel': 'cancel',
		'click .delete-mhsw': 'delete'
	},
	edit: function() {
		$('.edit-mhsw').hide();
		$('.delete-mhsw').hide();
		this.$('.update-mhsw').show();
		this.$('.cancel').show();

		var name = this.$('.name').html();
		var address = this.$('.address').html();
		var old = this.$('.old').html();

		this.$('.name').html('<input type="text" class="form-control name-update" value="' + name + '">');
		this.$('.address').html('<input type="text" class="form-control address-update" value="' + address + '">');
		this.$('.old').html('<input type="text" class="form-control old-update" value="' + old + '">');
	},
	update: function() {
		this.model.set('name', $('.name-update').val());
		this.model.set('address', $('.address-update').val());
		this.model.set('old', $('.old-update').val());

		this.model.save(null, {
			success: function(response) {
				console.log('Successfully UPDATED mhsw with _id: ' + response.toJSON()._id);
			},
			error: function(err) {
				console.log('Failed to update mhsw!');
			}
		});
	},
	cancel: function() {
		mhswsView.render();
	},
	delete: function() {
		this.model.destroy({
			success: function(response) {
				console.log('Successfully DELETED mhsw with _id: ' + response.toJSON()._id);
			},
			error: function(err) {
				console.log('Failed to delete mhsw!');
			}
		});
	},
	render: function() {
		//this.$el.html(this.template(this.model.attributes.data[0]))
		this.$el.html(this.template(this.model.toJSON()));
		return this;
	}
});

// Backbone View for all blogs

var MhswsView = Backbone.View.extend({
	model: mhsws,
	el: $('.mhsw-list'),
	initialize: function() {
		var self = this;
		this.model.on('add', this.render, this);
		this.model.on('change', function() {
			setTimeout(function() {
				self.render();
			}, 30);
		},this);
		this.model.on('remove', this.render, this);

		this.model.fetch({
			success: function(response) {
				_.each(response.toJSON(), function(item) {
					console.log('Successfully GOT mhsw with _id: ' + item._id);
				})
			},
			error: function() {
				console.log('Failed to get mhsw!');
			}
		});
	},
	render: function() {
		var self = this;
		this.$el.html('');
		_.each(this.model.toArray(), function(mhsw) {
			self.$el.append((new MhswView({model: mhsw})).render().$el);
		});
		return this;
	}
});

var mhswsView = new MhswsView();

$(document).ready(function() {
	$('.add-mhsw').on('click', function() {
		var mhsw = new Mhsw({
			//_id : '',
			name: $('.name-input').val(),
			address: $('.address-input').val(),
			old: $('.old-input').val()
		});
		$('.name-input').val('');
		$('.address-input').val('');
		$('.old-input').val('');
		mhsws.add(mhsw);
		mhsw.save(null, {
			success: function(response) {
				console.log('Successfully SAVED mhsw with _id: ' + response.toJSON()._id);
			},
			error: function(err) {
				console.log('Failed to save mhsw!'+err.toJSON()._id+','+err.toJSON().name+','+err.toJSON().address+','+err.toJSON().old);
				//console.log('JsonFile: '+err.toJSON());
			}
		});
	});
})
