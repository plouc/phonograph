var gulp       = require('gulp');
var browserify = require('browserify');
var reactify   = require('reactify');
var uglify     = require('gulp-uglify');
var source     = require('vinyl-source-stream');
var buffer     = require('vinyl-buffer');
var rename     = require('gulp-rename');
var watchify   = require('watchify');
var gutil      = require('gulp-util');
var chalk      = require('chalk');


function getBundler(isDev) {
    var bundler = browserify({
        entries:      ['./src/js/App.jsx'],
        extensions:   ['.js', '.jsx'],
        debug:        isDev,
        cache:        {},  // for watchify
        packageCache: {},  // for watchify
        fullPaths:    true // for watchify
    });

    bundler.transform(reactify, {
        es6: true
    });

    return bundler;
}

gulp.task('watch:js', function () {
    var watcher = watchify(getBundler(true));

    return watcher
        .on('error', gutil.log.bind(gutil, 'Browserify Error'))
        .on('update', function () {
            watcher.bundle()
                .pipe(source('vinyl.js'))
                .pipe(gulp.dest(config.dest))
                .pipe(buffer())
                .pipe(uglify())
                .pipe(rename({suffix: '.min'}))
                .pipe(gulp.dest('./dist'))
            ;

            gutil.log(chalk.green('Updated JavaScript sources'));
        })
        .bundle() // Create the initial bundle when starting the task
        .pipe(source('vinyl.js'))
        .pipe(gulp.dest('./dist'))
        .pipe(buffer())
        .pipe(uglify())
        .pipe(rename({suffix: '.min'}))
        .pipe(gulp.dest('./dist'))
        ;
});

gulp.task('watch:js:dev', function () {
    var watcher = watchify(getBundler(true));

    return watcher
        .on('error', gutil.log.bind(gutil, 'Browserify Error'))
        .on('update', function () {
            watcher.bundle()
                .pipe(source('vinyl.js'))
                .pipe(buffer())
                .pipe(gulp.dest('./dist'))
            ;

            gutil.log(chalk.green('Updated JavaScript sources [dev]'));
        })
        .bundle() // Create the initial bundle when starting the task
        .pipe(source('vinyl.js'))
        .pipe(gulp.dest('./dist'))
    ;
});


gulp.task('js:dev', function () {
    return getBundler(true)
        .bundle()
        .pipe(source('vinyl.js'))
        .pipe(gulp.dest('./dist'))
    ;
});


gulp.task('js', ['js:dev'], function () {
    return gulp.src('./dist/vinyl.js')
        .pipe(uglify())
        .pipe(rename({suffix: '.min'}))
        .pipe(gulp.dest('./dist'))
    ;
});