package com.fancyinnovations.fancyspaces.versions;

import com.fancyinnovations.fancyspaces.FancySpaces;
import com.fancyinnovations.fancyspaces.utils.HttpRequest;
import de.oliver.fancyanalytics.logger.properties.NumberProperty;
import de.oliver.fancyanalytics.logger.properties.StringProperty;
import de.oliver.fancyanalytics.logger.properties.ThrowableProperty;

import java.io.IOException;
import java.net.URISyntaxException;
import java.net.http.HttpResponse;
import java.util.Arrays;
import java.util.List;

public class VersionService {

    private final FancySpaces fs;

    public VersionService(FancySpaces fs) {
        this.fs = fs;
    }

    public List<Version> getVersions(String spaceID) {
        HttpRequest req = new HttpRequest(fs.getBaseURL() + "/spaces/" + spaceID + "/versions")
                .withHeader("Accept", "application/json")
                .withHeader("User-Agent", "FancySpaces Java-SDK");

        try {
            HttpResponse<String> resp = req.send();

            if (resp.statusCode() != 200) {
                fs.getFancyLogger().error(
                        "Failed to fetch latest version",
                        NumberProperty.of("status_code", resp.statusCode()),
                        StringProperty.of("response_body", resp.body())
                );
                return null;
            }

            Version[] fromJson = HttpRequest.gson.fromJson(resp.body(), Version[].class);
            return Arrays.stream(fromJson).toList();
        } catch (URISyntaxException | IOException | InterruptedException e) {
            fs.getFancyLogger().error(
                    "Exception occurred while fetching latest version",
                    ThrowableProperty.of(e)
            );
            return null;
        }
    }

    public Version getVersion(String spaceID, String version) {
        HttpRequest req = new HttpRequest(fs.getBaseURL() + "/spaces/" + spaceID + "/versions/"+version)
                .withHeader("Accept", "application/json")
                .withHeader("User-Agent", "FancySpaces Java-SDK");

        try {
            HttpResponse<String> resp = req.send();

            if (resp.statusCode() != 200) {
                fs.getFancyLogger().error(
                        "Failed to fetch latest version",
                        NumberProperty.of("status_code", resp.statusCode()),
                        StringProperty.of("response_body", resp.body())
                );
                return null;
            }

            return HttpRequest.gson.fromJson(resp.body(), Version.class);
        } catch (URISyntaxException | IOException | InterruptedException e) {
            fs.getFancyLogger().error(
                    "Exception occurred while fetching latest version",
                    ThrowableProperty.of(e)
            );
            return null;
        }
    }

    public Version getLatestVersion(String spaceID) {
        return getVersion(spaceID, "latest");
    }

}
