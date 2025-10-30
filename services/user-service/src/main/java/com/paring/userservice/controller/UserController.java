package com.paring.userservice.controller;

import com.paring.userservice.config.JwtUtil;
import com.paring.userservice.dto.LoginRequest;
import com.paring.userservice.dto.RegisterRequest;
import com.paring.userservice.dto.UserResponse;
import com.paring.userservice.model.User;
import com.paring.userservice.repository.UserRepository;
import com.paring.userservice.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.Authentication;
import org.springframework.web.bind.annotation.*;

import java.util.Map;

@RestController
@RequestMapping("/api/users")
public class UserController {

    @Autowired
    private UserService userService;

    private final JwtUtil jwtUtil;
    private final UserRepository userRepository;

    public UserController(UserService _userService, JwtUtil _jwtUtil, UserRepository _userRepository) {
        this.userService = _userService;
        this.jwtUtil = _jwtUtil;
        this.userRepository = _userRepository;
    }

    @PostMapping("/register")
    public ResponseEntity<UserResponse> register(@RequestBody RegisterRequest request) {
        return ResponseEntity.ok(userService.register(request));
    }

    @PostMapping("/login")
    public ResponseEntity<Map<String, String>> login(@RequestBody LoginRequest request) {
        return ResponseEntity.ok(userService.login(request));
    }

    @GetMapping("/profile")
    public ResponseEntity<UserResponse> getProfile(Authentication authentication) {
        String email = authentication.getName();
        System.out.println("[Controller] Email from token: " + email);
        return ResponseEntity.ok(userService.getProfile(email));
    }

    @PostMapping("/refresh")
    public ResponseEntity<Map<String, String>> refresh(@RequestHeader("Authorization") String refreshHeader) {
        if (refreshHeader == null || !refreshHeader.startsWith("Bearer ")) {
            throw new RuntimeException("Invalid refresh token");
        }

        String refreshToken = refreshHeader.substring(7);

        if (!jwtUtil.validateToken(refreshToken) || jwtUtil.isTokenExpired(refreshToken)) {
            throw new RuntimeException("Invalid or expired refresh token");
        }

        String email = jwtUtil.extractEmail(refreshToken);

        // Ambil user untuk generate ulang access token
        User user = userRepository.findByEmail(email)
                .orElseThrow(() -> new RuntimeException("User not found"));

        Map<String, String> newTokens = jwtUtil.generateTokens(email, user.getRole(), user.getId());

        return ResponseEntity.ok(newTokens);
    }
}