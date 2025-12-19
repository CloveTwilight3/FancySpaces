package com.fancyinnovations.fancyspaces;

import com.fancyinnovations.fancyspaces.versions.VersionService;
import de.oliver.fancyanalytics.logger.ExtendedFancyLogger;

public class FancySpaces {

    private final String baseURL;
    private final ExtendedFancyLogger fancyLogger;

    private final VersionService versionService;

    public FancySpaces() {
        this.baseURL = "https://fancyspaces.net/api/v1";
        this.fancyLogger = new ExtendedFancyLogger("FancySpaces Java-SDK");

        this.versionService = new VersionService(this);
    }

    public String getBaseURL() {
        return baseURL;
    }

    public ExtendedFancyLogger getFancyLogger() {
        return fancyLogger;
    }

    public VersionService getVersionService() {
        return versionService;
    }
}
