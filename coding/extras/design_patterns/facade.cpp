#include <iostream>
#include <string>

// Subsystem 1: Projector
class Projector {
public:
    void on() {
        std::cout << "Projector is ON.\n";
    }
    void off() {
        std::cout << "Projector is OFF.\n";
    }
    void setInput(const std::string& input) {
        std::cout << "Projector input set to: " << input << "\n";
    }
};

// Subsystem 2: Amplifier
class Amplifier {
public:
    void on() {
        std::cout << "Amplifier is ON.\n";
    }
    void off() {
        std::cout << "Amplifier is OFF.\n";
    }
    void setVolume(int level) {
        std::cout << "Amplifier volume set to: " << level << "\n";
    }
};

// Subsystem 3: DVD Player
class DVDPlayer {
public:
    void on() {
        std::cout << "DVD Player is ON.\n";
    }
    void off() {
        std::cout << "DVD Player is OFF.\n";
    }
    void play(const std::string& movie) {
        std::cout << "Playing movie: " << movie << "\n";
    }
    void stop() {
        std::cout << "DVD Player stopped.\n";
    }
};

// Facade: HomeTheaterFacade
class HomeTheaterFacade {
public:
    HomeTheaterFacade(Projector& projector, Amplifier& amplifier, DVDPlayer& dvdPlayer)
        : projector_(projector), amplifier_(amplifier), dvdPlayer_(dvdPlayer) {}

    void watchMovie(const std::string& movie) {
        std::cout << "\nGetting ready to watch a movie...\n";
        projector_.on();
        projector_.setInput("DVD");
        amplifier_.on();
        amplifier_.setVolume(10);
        dvdPlayer_.on();
        dvdPlayer_.play(movie);
        std::cout << "Movie setup is complete. Enjoy the movie!\n";
    }

    void endMovie() {
        std::cout << "\nShutting down movie theater...\n";
        dvdPlayer_.stop();
        dvdPlayer_.off();
        amplifier_.off();
        projector_.off();
        std::cout << "Movie theater is shut down.\n";
    }

private:
    Projector& projector_;
    Amplifier& amplifier_;
    DVDPlayer& dvdPlayer_;
};

// Main Function (Driver Code)
int main() {
    // Create subsystem components
    Projector projector;
    Amplifier amplifier;
    DVDPlayer dvdPlayer;

    // Create a facade for the home theater
    HomeTheaterFacade homeTheater(projector, amplifier, dvdPlayer);

    // Use the facade to control the subsystems
    homeTheater.watchMovie("Inception");
    homeTheater.endMovie();

    return 0;
}
