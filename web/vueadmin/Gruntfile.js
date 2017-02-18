'use strict';

const HtmlWebpackPlugin = require('html-webpack-plugin');
 
module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    webpack: {
      dist: {
        entry: './src/main.js',
        output: {
          filename: 'main.js',
          path: 'dist'
        },
        stats: { colors: true },
        module: {
          rules: [
            { test: /\.vue$/, loader: 'vue-loader' }
          ]
        },
        resolve: {
          alias: {
            'vue$': 'vue/dist/vue.common.js'
          }
        },
        plugins: [
          new HtmlWebpackPlugin({
            title: 'Spruce CMS Application',
            filename: 'index.html',
            hash: true,
            template: "src/index.html"
          })
        ]
      }
    },

/*
    uglify: {
      dist: {
        options: {
          sourceMap: true,
          compress: false
        },
        files: {
          'js/lib.js': [
            'js/lib/jquery-2.1.3.min.js',
            'js/lib/underscore-min.js',
            'js/lib/angular.min.js',
            'js/lib/jquery.hoverIntent-1.8.0.js',
            'js/lib/fastclick.js',
            'js/lib/presentation.js',
            'js/lib/matchHeight.js',
            'js/lib/navigation.js',
            'js/lib/jquery.fitvids.js'
          ],
          'js/app.js': [
            'js/app/main.js',
          ],
        }
      }
    },
  */

    watch: {
      options: {
        livereload: true
      },
      js: {
        files: ['src/**/*.js', 'src/**/*.html', 'src/**/*.vue'],
        tasks: [
        'jshint', 
        'webpack']
      },
      css: {
        files: ['src/**/*.scss'],
        tasks: ['sass']
      }
    },

    jshint: {
      all: {
        options: {
          esversion: 6
        },
        src: [ 'src/**/*.js' ]
      }
    },

    sass: {
      dist: {
        options: {
          style: 'compressed'
        },
        files: {
          'dist/main.css': 'src/main.scss',
        }
      }
    },


  });

  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-webpack');
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-jshint');

  grunt.registerTask('watchcss', ['watch:css']);
  grunt.registerTask('watchsass', ['watch:css']);
  grunt.registerTask('watchjs', ['watch:js']);

  grunt.registerTask('default', [ 'jshint', 'webpack', 'sass' ]);
  grunt.registerTask('build', ['jshint', 'webpack', 'sass' ]);
};
