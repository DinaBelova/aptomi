<template>
  <div>

    <!-- /.row -->
    <div class="row">
      <div class="col-xs-12">
        <div class="box">
          <div class="box-header">
            <h3 class="box-title">Declared Dependencies</h3>
          </div>
          <!-- /.box-header -->
          <div class="overlay" v-if="loading">
            <i class="fa fa-refresh fa-spin"></i>
          </div>
          <div class="box-body table-responsive no-padding">
            <table class="table table-hover">
              <thead>
              <tr>
                <th>Namespace</th>
                <th>Name</th>
                <th>User</th>
                <th>Contract</th>
                <th>Status</th>
                <th>Action</th>
              </tr>
              </thead>
              <tbody>
              <tr v-if="error">
                <td><span class="label label-danger center">Error</span> <i class="text-red">{{ error }}</i></td>
              </tr>
              <tr v-if="dependencies == null || dependencies.length <= 0">
                <td>No Dependencies Declared</td>
              </tr>
              <tr v-for="d in dependencies">
                <td>{{d.namespace}}</td>
                <td>{{d.name}}</td>
                <td v-if="!d.error">{{d.user}}</td>
                <td v-else><span class="label label-danger center">Error</span></td>
                <td v-if="!d.error">{{d.contract}}</td>
                <td v-else><span class="label label-danger center">Error</span></td>
                <td v-if="!d.status_error">
                  <span class="label" v-bind:class="{ 'label-success': d['status'] === 'Deployed', 'label-warning': d['status'] !== 'Deployed'}">{{d.status}}</span>
                </td>
                <td v-else><span class="label label-danger center">Error</span></td>
                <td>
                  <button v-if="d['status'] === 'Deployed'" type="button" class="btn btn-default btn-xs" @click="showEndpointsForDependency = d">Show
                    Endpoints
                  </button>
                </td>
              </tr>
              </tbody>
            </table>
          </div>
          <!-- /.box-body -->
        </div>
        <!-- /.box -->
      </div>
    </div>

    <!-- /.row -->
    <div class="row" v-if="showEndpointsForDependency">
      <div class="col-xs-12">
        <endpoints :dependency="showEndpointsForDependency"></endpoints>
      </div>
    </div>

  </div>
</template>

<script>
  import {getDependencies} from 'lib/api.js'
  import Endpoints from 'pages/components/Endpoints'

  export default {
    data () {
      // empty data
      return {
        loading: false,
        dependencies: null,
        error: null,
        showEndpointsForDependency: null
      }
    },
    created () {
      // fetch the data when the view is created and the data is already being observed
      this.fetchData()
    },
    components: {
      'endpoints': Endpoints
    },
    methods: {
      fetchData () {
        this.loading = true
        this.dependencies = null
        this.error = null

        const fetchSuccess = $.proxy(function (data) {
          this.loading = false
          this.dependencies = data
        }, this)

        const fetchError = $.proxy(function (err) {
          this.loading = false
          this.error = err
        }, this)

        getDependencies(fetchSuccess, fetchError)
      }
    }
  }
</script>

<style>

</style>
