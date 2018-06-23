const gulp = require('gulp')
const wait = require('gulp-wait')
const sass = require('gulp-sass')
const prefixer = require('gulp-autoprefixer')
const cssnano = require('gulp-cssnano')

gulp.task('default', ['scss'])

gulp.task('scss', () => {
    return gulp.src('scss/*.scss')
        .pipe(wait(200))
        .pipe(sass().on('error', sass.logError))
        .pipe(prefixer())
        .pipe(cssnano())
        .pipe(gulp.dest('dist/'))
})

gulp.task('watch', () => {
    gulp.watch('scss/**/*.scss', ['scss'])
})