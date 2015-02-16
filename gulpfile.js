var gulp = require('gulp');
var shell = require('gulp-shell');
var watch = require('gulp-watch');

gulp.task('proxy', shell.task([
    'screen -d -m browser-sync start --files "default/*.html, default/*.css, default/*.js" --proxy="localhost:8080"'
]));

gulp.task('kill', shell.task([
    'screen -ls | grep Detached | cut -d. -f1 | awk "{print $1}" | xargs kill'
]));

gulp.task('default', shell.task([
    '~/appengine/dev_appserver.py default.yaml'
]));

gulp.task('build', shell.task([
    'vulcanize -o build/index.html default/index.html --inline --strip'
]));

gulp.task('watch', function() {
    gulp.watch('default/**.*', ['build']);
});

gulp.task('appserver', ['build', 'watch'], shell.task([
    '~/appengine/dev_appserver.py build.yaml'
]));

gulp.task('update', ['build'], shell.task([
    '~/appengine/appcfg.py --oauth2 update build.yaml'
]));

gulp.task('dispatch', shell.task([
    '~/appengine/appcfg.py --oauth2 update_dispatch dispatch.yaml build.yaml'
]));

gulp.task('index', shell.task([
    '~/appengine/appcfg.py --oauth2 update_indexes index.yaml build.yaml'
]));

gulp.task('vacuum', shell.task([
    '~/appengine/appcfg.py --oauth2 vacuum_indexes index.yaml build.yaml'
]));

gulp.task('rollback', shell.task([
    '~/appengine/appcfg.py --oauth2 rollback build.yaml build.yaml'
]));
