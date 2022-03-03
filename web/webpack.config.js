const Encore = require('@symfony/webpack-encore')

if (!Encore.isRuntimeEnvironmentConfigured()) {
    Encore.configureRuntimeEnvironment(process.env.NODE_ENV || 'dev');
}

Encore
    .setOutputPath('public/build')
    .setPublicPath('/build')

    .enableVersioning(false) // until we do


    .disableSingleRuntimeChunk()
    .cleanupOutputBeforeBuild()

    .addEntry('app', './assets/js/app.js')

    .enableSassLoader()

    .configureBabelPresetEnv((config) => {
        config.useBuiltIns = 'usage'
        config.corejs = 3
    })


module.exports = Encore.getWebpackConfig()